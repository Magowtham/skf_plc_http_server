package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type DeleteDrierUseCase struct {
	DataBaseService *service.DataBaseService
	CacheService    *service.CacheService
}

func InitDeleteDrierUseCase(dbRepo repository.DataBaseRepository, cacheRepo repository.CacheRepository) *DeleteDrierUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	cacheService := service.NewCacheService(cacheRepo)
	return &DeleteDrierUseCase{
		DataBaseService: dbService,
		CacheService:    cacheService,
	}
}

func (u *DeleteDrierUseCase) Execute(plcId string, drierId string) (error, int) {

	if plcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1
	}

	if drierId == "" {
		return fmt.Errorf("drier id cannot be empty"), 1
	}

	isPlcIdExists, error := u.DataBaseService.CheckPlcIdExists(plcId)

	if error != nil {
		log.Printf("error occurred with database while checking plc id exists, plc id -> %s drier id -> %s\n", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isPlcIdExists {
		return fmt.Errorf("plc id not exists"), 1
	}

	isDrierIdExists, error := u.DataBaseService.CheckDrierIdExists(drierId)

	if error != nil {
		log.Printf("error occurred with database while checking drier id exists, plc id -> %s drier id -> %s\n", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isDrierIdExists {
		return fmt.Errorf("drier id not exists"), 1
	}

	regAddresses, err := u.DataBaseService.GetRegisterAddressesByDrierId(plcId, drierId)

	if err != nil {
		log.Printf("error occurred with database while getting the register addresses, plc id -> %s, drier id -> %s\n", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2
	}

	for _, regAddress := range regAddresses {
		if err := u.CacheService.DeleteRegister(plcId, regAddress); err != nil {
			log.Printf("error occurred with redis while deleting the register, plc id -> %s, drier id -> %s", plcId, drierId)
			return fmt.Errorf("error occurred with cache"), 2
		}
	}

	if err := u.CacheService.DeleteDrier(drierId); err != nil {
		log.Printf("error occurrred with redis while deleting the drier plc id -> %s, drier id -> %s", plcId, drierId)
		return fmt.Errorf("error occurred with cache"), 2
	}

	if error := u.DataBaseService.DeleteDrier(drierId); error != nil {
		log.Printf("error occurred with database while deleting the drier, plc id -> %s, drier id -> %s\n", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
