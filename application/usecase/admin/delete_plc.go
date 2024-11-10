package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type DeletePlcUseCase struct {
	DatabaseService *service.DataBaseService
	CacheService    *service.CacheService
}

func InitDeletePlcUseCase(dbRepo repository.DataBaseRepository, cacheRepo repository.CacheRepository) *DeletePlcUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	cacheService := service.NewCacheService(cacheRepo)
	return &DeletePlcUseCase{
		DatabaseService: dbService,
		CacheService:    cacheService,
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

	regAddresses, err := u.DatabaseService.GetAllRegisterAddress(plcId)

	if err != nil {
		log.Printf("error occurred with database while getting all register address, delete plc, plc id -> %s", plcId)
		return fmt.Errorf("error occurred with database"), 2
	}

	for _, regAddress := range regAddresses {
		if err := u.CacheService.DeleteRegister(plcId, regAddress); err != nil {
			log.Printf("error occurred with redis while deleting the register, delete plc, plc id -> %s", plcId)
			return fmt.Errorf("error occurred with database"), 2
		}
	}

	driers, err := u.DatabaseService.GetDriersByPlcId(plcId)

	if err != nil {
		log.Printf("error occurred with database while getting driers, plc id -> %s", plcId)
		return fmt.Errorf("error occurred with database"), 2
	}

	for _, drier := range driers {
		if err := u.CacheService.DeleteDrier(drier.DrierId); err != nil {
			log.Printf("error occurred with database while deleting drier, delete plc, plc id -> %s", plcId)
			return fmt.Errorf("error occurred with database"), 2
		}
	}

	if error := u.DatabaseService.DeletePlc(plcId); error != nil {
		log.Printf("error occurred with database while deleting the plc having plc id -> %s", plcId)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
