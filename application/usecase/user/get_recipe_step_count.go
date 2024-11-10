package user

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type GetRecipeStepCountUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitGetRecipeStepCountUseCase(dbRepo repository.DataBaseRepository) *GetRecipeStepCountUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &GetRecipeStepCountUseCase{
		DataBaseService: dbService,
	}
}

func (u *GetRecipeStepCountUseCase) Execute(drierId string) (error, int, int) {
	if drierId == "" {
		return fmt.Errorf("drier id cannot be empty"), 1, 0
	}

	isDrierIdExists, error := u.DataBaseService.CheckDrierIdExists(drierId)

	if error != nil {
		log.Printf("error occurred with database while checking drier id exists, get recipe step count, drier id -> %s", drierId)
		return fmt.Errorf("error occurred with database"), 2, 0
	}

	if !isDrierIdExists {
		return fmt.Errorf("drier id not exists"), 1, 0
	}

	recipeStepCount, error := u.DataBaseService.GetRecipeStepCount(drierId)

	if error != nil {
		log.Printf("error occurred with database while checking getting recipe step count, get recipe step count, drier id -> %s", drierId)
		return fmt.Errorf("error occurred with database"), 2, 0
	}

	return nil, 0, recipeStepCount
}
