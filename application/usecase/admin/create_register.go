package admin

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
)

type CreateRegisterUseCase struct {
	DataBaseService *service.DataBaseService
	CacheService    *service.CacheService
}

func InitCreateRegisterUseCase(dbRepo repository.DataBaseRepository, cacheRepo repository.CacheRepository) *CreateRegisterUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	cacheService := service.NewCacheService(cacheRepo)
	return &CreateRegisterUseCase{
		DataBaseService: dbService,
		CacheService:    cacheService,
	}
}

func (u *CreateRegisterUseCase) Execute(plcId string, drierId string, registerRequest *request.Register) (error, int) {

	if plcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1
	}

	if drierId == "" {
		return fmt.Errorf("drier id cannot be empty"), 1
	}

	if registerRequest.RegAddress == "" {
		return fmt.Errorf("register address cannot be empty"), 1
	}

	if registerRequest.RegType == "" {
		return fmt.Errorf("register type cannot be empty"), 1
	}

	if registerRequest.Label == "" {
		return fmt.Errorf("label cannot be empty"), 1
	}

	isPlcIdExists, error := u.DataBaseService.CheckPlcIdExists(plcId)

	if error != nil {
		log.Printf("error occurred with database while checking plc id exists, plc id -> %s", plcId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isPlcIdExists {
		return fmt.Errorf("plc id not exists"), 1
	}

	isDrierIdExists, error := u.DataBaseService.CheckDrierIdExists(drierId)

	if error != nil {
		log.Printf("error occurred with database while checking direr id exists, drier id -> %s", drierId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isDrierIdExists {
		return fmt.Errorf("drier id not exists"), 2
	}

	isRegisterAddressExists, error := u.DataBaseService.CheckRegisterAddressExists(plcId, registerRequest.RegAddress)

	if error != nil {
		log.Printf("error occurred while checking register address exists, plc id -> %s , direr id -> %s , register address -> %s ", plcId, drierId, registerRequest.RegAddress)
		return fmt.Errorf("error occurred with database"), 2
	}

	if isRegisterAddressExists {
		return fmt.Errorf("register address exist"), 1
	}

	isRegisterTypeExists, error := u.DataBaseService.CheckRegisterTypeExists(plcId, drierId, registerRequest.RegType)

	if error != nil {
		log.Printf("error occurred with database while checking register type exists, plc id -> %s, drier id -> %s, register address -> %s, register type -> %s", plcId, drierId, registerRequest.RegAddress, registerRequest.RegType)
		return fmt.Errorf("error occurred with database"), 2
	}

	if isRegisterTypeExists {
		return fmt.Errorf("register type exists"), 1
	}

	regex := regexp.MustCompile(`^stptmp\d+$`)

	register := &entity.Register{
		DrierId:             drierId,
		RegAddress:          registerRequest.RegAddress,
		RegType:             registerRequest.RegType,
		Label:               registerRequest.Label,
		Value:               "0",
		LastUpdateTimestamp: time.Now().UTC(),
	}

	if regex.MatchString(registerRequest.RegType) {
		if error := u.DataBaseService.UpdateDrierRecipeStepCountAndCreateRegister(plcId, register); error != nil {
			log.Printf("error occurred while update and creating the register, create register, plc id -> %s, drier id -> %s", plcId, drierId)
			return fmt.Errorf("error occurred with database"), 2
		}

		return nil, 0
	}

	if error := u.DataBaseService.CreateRegister(plcId, register); error != nil {
		log.Printf("error occurred with database while creating the register having plc id -> %s and drier id -> %s\n", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if error := u.CacheService.CreateRegister(plcId, register); error != nil {
		log.Printf("error occurred with cache while creating the register having plc id -> %s and drier id -> %s\n", plcId, drierId)
		return fmt.Errorf("error occurred with cache"), 2
	}

	return nil, 0
}
