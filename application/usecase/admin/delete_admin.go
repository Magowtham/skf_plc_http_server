package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type DeleteAdminUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitDeleteAdminUaseCase(repo repository.DataBaseRepository) DeleteAdminUseCase {
	service := service.NewDataBaseService(repo)
	return DeleteAdminUseCase{
		DataBaseService: service,
	}
}

func (u *DeleteAdminUseCase) Execute(adminId string) (error, int) {
	if adminId == "" {
		return fmt.Errorf("admin id cannot be empty"), 1
	}

	isAdminIdExists, error := u.DataBaseService.CheckAdminIdExists(adminId)

	if error != nil {
		log.Printf("error occurred with database while checking admin id -> %s exists\n", adminId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isAdminIdExists {
		return fmt.Errorf("invalid admin id"), 1
	}

	error = u.DataBaseService.DeleteAdmin(adminId)

	if error != nil {
		log.Printf("error occurred with database while deleting the admin id -> %s\n", adminId)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0

}
