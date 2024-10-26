package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
)

type CreatePlcUseCase struct {
	DatabaseService service.DataBaseService
}

func InitCreatePlcUseCase(repo repository.DataBaseRepository) CreatePlcUseCase {
	service := service.NewDataBaseService(repo)
	return CreatePlcUseCase{
		DatabaseService: service,
	}
}

func (u *CreatePlcUseCase) Execute(userId string, plcRequest request.Plc) (error, int) {
	if userId == "" {
		return fmt.Errorf("user id cannot be empty"), 1
	}

	if plcRequest.Label == "" {
		return fmt.Errorf("label cannot be empty"), 1
	}

	if plcRequest.PlcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1
	}

	isUserIdExists, error := u.DatabaseService.CheckUserIdExists(userId)

	if error != nil {
		log.Printf("error occurred with database while checking user user id exists for user id -> %s and plc id -> %s", userId, plcRequest.PlcId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isUserIdExists {
		return fmt.Errorf("user id not exists"), 1
	}

	isPlcIdExists, error := u.DatabaseService.CheckPlcIdExists(plcRequest.PlcId)

	if error != nil {
		log.Printf("error occurred with database while checking plc id exists for user id -> %s and plc id -> %s", userId, plcRequest.PlcId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if isPlcIdExists {
		return fmt.Errorf("plc id already exists"), 1
	}

	plc := entity.Plc{
		PlcId:  plcRequest.PlcId,
		UserId: userId,
		Label:  plcRequest.Label,
	}

	error = u.DatabaseService.CreatePlc(plc)

	if error != nil {
		log.Printf("error occurred with database while creating the plc having user id -> %s plc id -> %s", userId, plcRequest.PlcId)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
