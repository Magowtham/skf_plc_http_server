package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type DeleteRegTypeUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitDeleteRegTypeUseCase(repo repository.DataBaseRepository) *DeleteRegTypeUseCase {
	service := service.NewDataBaseService(repo)

	return &DeleteRegTypeUseCase{
		DataBaseService: service,
	}
}

func (u *DeleteRegTypeUseCase) Execute(regTypeName string) (error, int) {
	if regTypeName == "" {
		return fmt.Errorf("register type cannot be empty"), 1
	}
	isRegTypeNameExists, error := u.DataBaseService.CheckRegTypeNameExistsInRegTypes(regTypeName)

	if error != nil {
		log.Printf("error occurred with database while checking reg type exists in reg types, delete reg type, reg type -> %s", regTypeName)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isRegTypeNameExists {
		return fmt.Errorf("register type not exists"), 1
	}

	if error := u.DataBaseService.DeleteRegType(regTypeName); error != nil {
		log.Printf("error occurred with database while deleting the reg type, delete reg type, reg type -> %s", regTypeName)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
