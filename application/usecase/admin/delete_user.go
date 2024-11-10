package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type DeleteUserUseCase struct {
	DataBaseService *service.DataBaseService
	CacheService    *service.CacheService
}

func InitDeleteUserCase(dbRepo repository.DataBaseRepository, cacheRepo repository.CacheRepository) *DeleteUserUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	cacheService := service.NewCacheService(cacheRepo)

	return &DeleteUserUseCase{
		DataBaseService: dbService,
		CacheService:    cacheService,
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

	plcs, err := u.DataBaseService.GetPlcsByUserId(userId)

	if err != nil {
		log.Printf("error occurrred while getting plcs by user id, delete user, user id -> %s", userId)
		return fmt.Errorf("error occurred with database"), 2
	}

	for _, plc := range plcs {
		regAddresses, err := u.DataBaseService.GetAllRegisterAddress(plc.PlcId)

		if err != nil {
			log.Printf("error occurred with database while getting register addresses, delete user, user id -> %s plc id -> %s", userId, plc.PlcId)
			return fmt.Errorf("error occurred with database"), 2
		}

		for _, regAddress := range regAddresses {
			if err := u.CacheService.DeleteRegister(plc.PlcId, regAddress); err != nil {
				log.Printf("error occurred with redis while deleting the register, delete user, user id -> %s, plc id -> %s", userId, plc.PlcId)
				return fmt.Errorf("error occurred with cache"), 2
			}
		}

		driers, err := u.DataBaseService.GetDriersByPlcId(plc.PlcId)

		if err != nil {
			log.Printf("error occurred with database while getting driers by plc id,delete user, user id -> %s , plc id -> %s", userId, plc.PlcId)
			return fmt.Errorf("error occurred with database"), 2
		}

		for _, drier := range driers {
			if err := u.CacheService.DeleteDrier(drier.DrierId); err != nil {
				log.Printf("error occurred with redis while deleting the drier, delete user, user id -> %s, plc id -> %s, drier id -> %s", userId, plc.PlcId, drier.DrierId)
				return fmt.Errorf("error occurred with cache"), 2
			}
		}
	}

	error = u.DataBaseService.DeleteUser(userId)

	if error != nil {
		log.Printf("error occurred with database while deleting the user -> %s", userId)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
