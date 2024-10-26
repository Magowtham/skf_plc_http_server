package admin

import (
	"fmt"
	"log"
	"regexp"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type DeleteRegisterUseCase struct {
	DataBaseService service.DataBaseService
}

func InitDeleteRegisterUseCase(repo repository.DataBaseRepository) DeleteRegisterUseCase {
	service := service.NewDataBaseService(repo)

	return DeleteRegisterUseCase{
		DataBaseService: service,
	}
}

func (u *DeleteRegisterUseCase) Execute(plcId string, drierId string, registerAddress string, registerType string) (error, int) {
	if plcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1
	}

	if drierId == "" {
		return fmt.Errorf("drier id cannot be empty"), 1
	}

	if registerAddress == "" {
		return fmt.Errorf("register address cannot be empty"), 1
	}

	if registerType == "" {
		return fmt.Errorf("register type cannot be empty"), 1
	}

	isPlcIdExists, error := u.DataBaseService.CheckPlcIdExists(plcId)

	if error != nil {
		log.Printf("error occurred with database while checking plc id exists, delete register, plc id -> %s, drier id -> %s reg address -> %s reg type -> %s ", plcId, drierId, registerAddress, registerType)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isPlcIdExists {
		return fmt.Errorf("plc id not exists"), 1
	}

	isRegisterAddressExists, error := u.DataBaseService.CheckRegisterAddressAndRegisterTypeExists(plcId, registerAddress, registerType)

	if error != nil {
		log.Printf("error occurred with database while checking register address exists, delete register, plc id -> %s, drier id -> %s reg address -> %s reg type -> %s ", plcId, drierId, registerAddress, registerType)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isRegisterAddressExists {
		return fmt.Errorf("register address not exists"), 1
	}

	regex := regexp.MustCompile(`^stptmp\d+$`)

	if regex.MatchString(registerType) {
		if error := u.DataBaseService.UpdateDrierRecipeStepCountAndDeleteRegisterByRegAddress(plcId, drierId, registerAddress); error != nil {
			log.Printf("error occurred with database while update and deleting register, plc id -> %s, drier id -> %s reg address -> %s reg type -> %s ", plcId, drierId, registerAddress, registerType)
			return fmt.Errorf("error occurred with database"), 2
		}
		return nil, 0
	}

	if error := u.DataBaseService.DeleteRegisterByRegAddress(plcId, registerAddress); error != nil {
		log.Printf("error occurred with database while deleting register, delete register, plc id -> %s, drier id -> %s, reg address -> %s, reg type -> %s", plcId, drierId, registerAddress, registerType)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
