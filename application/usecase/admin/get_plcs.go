package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type GetPlcsUseCase struct {
	DataBaseService service.DataBaseService
}

func InitGetPlcsUseCase(repo repository.DataBaseRepository) GetPlcsUseCase {
	service := service.NewDataBaseService(repo)
	return GetPlcsUseCase{
		DataBaseService: service,
	}
}

func (u *GetPlcsUseCase) Execute(userId string) (error, int, []entity.Plc) {
	if userId == "" {
		return fmt.Errorf("user id cannot be empty"), 1, nil
	}

	isUserIdExists, error := u.DataBaseService.CheckUserIdExists(userId)

	if error != nil {
		log.Printf("error occurred with database while checking user id exists in get user having user id -> %s\n", userId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isUserIdExists {
		return fmt.Errorf("user id not exists"), 1, nil
	}

	plcs, error := u.DataBaseService.GetPlcsByUserId(userId)

	if error != nil {
		log.Printf("error occurred with database while getting plcs. user id -> %s", userId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	return nil, 0, plcs
}
