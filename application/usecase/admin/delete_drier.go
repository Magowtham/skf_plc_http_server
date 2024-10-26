package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type DeleteDrierUseCase struct {
	DataBaseService service.DataBaseService
}

func InitDeleteDrierUseCase(repo repository.DataBaseRepository) DeleteDrierUseCase {
	service := service.NewDataBaseService(repo)
	return DeleteDrierUseCase{
		DataBaseService: service,
	}
}

func (u *DeleteDrierUseCase) Execute(drierId string) (error, int) {
	if drierId == "" {
		return fmt.Errorf("drier id cannot be empty"), 1
	}

	isDrierIdExists, error := u.DataBaseService.CheckDrierIdExists(drierId)

	if error != nil {
		log.Printf("error occurred with database while checking drier id exists, drier id -> %s\n", drierId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isDrierIdExists {
		return fmt.Errorf("drier id not exists"), 1
	}

	if error := u.DataBaseService.DeleteDrier(drierId); error != nil {
		log.Printf("error occurred with database while deleting the drier, drier id -> %s\n", drierId)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
