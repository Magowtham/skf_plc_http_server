package user

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type GetDriersUseCase struct {
	DataBaseService service.DataBaseService
}

func InitGetDriersUseCase(dbRepo repository.DataBaseRepository) GetDriersUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return GetDriersUseCase{
		DataBaseService: dbService,
	}
}

func (u *GetDriersUseCase) Execute(userId string) (error, int, []entity.Drier) {
	if userId == "" {
		return fmt.Errorf("user id cannot be empty"), 1, nil
	}

	isUserIdExists, error := u.DataBaseService.CheckUserIdExists(userId)

	if error != nil {
		log.Printf("error occurred with database while checking user id exists, get driers for users, user id ->%s", userId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isUserIdExists {
		return fmt.Errorf("user id not exists"), 1, nil
	}

	driers, error := u.DataBaseService.GetDriersByUserId(userId)

	if error != nil {
		log.Printf("error occurred with database while getting driers by user id, get driers for users, user id ->%s", userId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	return nil, 0, driers
}
