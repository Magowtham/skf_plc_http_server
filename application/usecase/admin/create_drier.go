package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
	"github.com/google/uuid"
)

type CreateDrierUseCase struct {
	DataBaseService service.DataBaseService
}

func InitCreateDrierUseCase(repo repository.DataBaseRepository) CreateDrierUseCase {
	service := service.NewDataBaseService(repo)
	return CreateDrierUseCase{
		DataBaseService: service,
	}
}

func (u *CreateDrierUseCase) Execute(plcId string, drierRequest *request.Drier) (error, int) {
	if plcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1
	}

	if drierRequest.Label == "" {
		return fmt.Errorf("label cannot be empty"), 1
	}

	isPlcIdExists, error := u.DataBaseService.CheckPlcIdExists(plcId)

	if error != nil {
		log.Printf("error occurred with database while checking plc id exists in creating drier having plc id -> %s drier label -> %s", plcId, drierRequest.Label)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isPlcIdExists {
		return fmt.Errorf("plc id not exists"), 1
	}

	uuid := uuid.New().String()
	drier := entity.Drier{
		DrierId:         uuid,
		PlcId:           plcId,
		RecipeStepCount: "0",
		Label:           drierRequest.Label,
	}

	if error := u.DataBaseService.CreateDrier(&drier); error != nil {
		log.Printf("error occurred with database while creating the drier having plcid -> %s and label -> %s\n", drier.PlcId, drier.Label)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
