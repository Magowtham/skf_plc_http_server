package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type GetDriersUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitGetDriersUseCase(repo repository.DataBaseRepository) GetDriersUseCase {
	service := service.NewDataBaseService(repo)

	return GetDriersUseCase{
		DataBaseService: service,
	}
}

func (u *GetDriersUseCase) Execute(plcId string) (error, int, []entity.Drier) {
	if plcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1, nil
	}

	isPlcIdExists, error := u.DataBaseService.CheckPlcIdExists(plcId)

	if error != nil {
		log.Printf("error occurred with database while checking plc id exists at a time of getting driers having plc id -> %s\n", plcId)
		return fmt.Errorf("error occurred with the database"), 2, nil
	}

	if !isPlcIdExists {
		return fmt.Errorf("plc id not exists"), 1, nil
	}

	driers, error := u.DataBaseService.GetDriersByPlcId(plcId)

	if error != nil {
		log.Printf("error occurred with database while getting the driers having plc id -> %s\n", plcId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	return nil, 0, driers
}
