package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type DeleteUserUseCase struct {
	DataBaseService service.DataBaseService
}

func InitDeleteUserCase(repo repository.DataBaseRepository) DeleteUserUseCase {
	service := service.NewDataBaseService(repo)

	return DeleteUserUseCase{
		DataBaseService: service,
	}
}

func (u *DeleteUserUseCase) Execute(userId string) (error, int) {
	if userId == "" {
		return fmt.Errorf("user id cannot be empty"), 1
	}

	isUserIdExists, error := u.DataBaseService.CheckUserIdExists(userId)

	if error != nil {
		log.Printf("error occurred with database while checking user id exists -> %s", userId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isUserIdExists {
		return fmt.Errorf("user not exists"), 1
	}

	error = u.DataBaseService.DeleteUser(userId)

	if error != nil {
		log.Printf("error occurred with database while deleting the user -> %s", userId)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
