package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type GetAllUsersUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitGetAllUsersUseCase(repo repository.DataBaseRepository) GetAllUsersUseCase {
	service := service.NewDataBaseService(repo)

	return GetAllUsersUseCase{
		DataBaseService: service,
	}
}

func (u *GetAllUsersUseCase) Execute() (error, int, []entity.User) {
	users, error := u.DataBaseService.GetAllUsers()

	if error != nil {
		log.Printf("error occurred with database while getting all the users")
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	return nil, 0, users
}
