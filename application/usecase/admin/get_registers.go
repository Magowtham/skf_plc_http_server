package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type GetRegistersUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitGetRegisterUseCase(repo repository.DataBaseRepository) *GetRegistersUseCase {
	service := service.NewDataBaseService(repo)
	return &GetRegistersUseCase{
		DataBaseService: service,
	}
}

func (u *GetRegistersUseCase) Execute(plcId string, drierId string) (error, int, []*entity.Register) {
	if plcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1, nil
	}

	if drierId == "" {
		return fmt.Errorf("drier id cannot be empty"), 1, nil
	}

	isPlcIdExists, error := u.DataBaseService.CheckPlcIdExists(plcId)

	if error != nil {
		log.Printf("error occurred with database while checking plc id exists, getting registers, plc id -> %s, direr id -> %s", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isPlcIdExists {
		return fmt.Errorf("plc id not exists"), 1, nil
	}

	isDrierIdExists, error := u.DataBaseService.CheckDrierIdExists(drierId)

	if error != nil {
		log.Printf("error occurred with database while checking drier id exists, getting registers, plc id -> %s, drier id -> %s", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isDrierIdExists {
		return fmt.Errorf("drier id not exists"), 1, nil
	}

	registers, error := u.DataBaseService.GetRegistersByDrierId(plcId, drierId)

	if error != nil {
		log.Printf("error occurred with database while getting driers, getting registers, plc id -> %s, drier id -> %s", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	return nil, 0, registers
}
