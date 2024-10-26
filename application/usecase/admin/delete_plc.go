package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type DeletePlcUseCase struct {
	DatabaseService service.DataBaseService
}

func InitDeletePlcUseCase(repo repository.DataBaseRepository) DeletePlcUseCase {
	service := service.NewDataBaseService(repo)
	return DeletePlcUseCase{
		DatabaseService: service,
	}
}

func (u *DeletePlcUseCase) Execute(plcId string) (error, int) {
	if plcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1
	}

	isPlcIdExists, error := u.DatabaseService.CheckPlcIdExists(plcId)

	if error != nil {
		log.Printf("error occurred with database while checking plc id exists in deleting the plc having plc id -> %s", plcId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isPlcIdExists {
		return fmt.Errorf("plc id not exists"), 1
	}

	if error := u.DatabaseService.DeletePlc(plcId); error != nil {
		log.Printf("error occurred with database while deleting the plc having plc id -> %s", plcId)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
