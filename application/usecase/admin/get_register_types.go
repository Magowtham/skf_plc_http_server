package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type GetRegisterTypesUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitGetRegisterTypesUseCase(dbRepo repository.DataBaseRepository) *GetRegisterTypesUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &GetRegisterTypesUseCase{
		DataBaseService: dbService,
	}
}

func (u *GetRegisterTypesUseCase) Execute(plcId string, drierId string) (error, int, []*entity.RegisterType) {

	if plcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1, nil
	}

	if drierId == "" {
		return fmt.Errorf("direr id cannot be empty"), 1, nil
	}

	isPlcIdExists, error := u.DataBaseService.CheckPlcIdExists(plcId)

	if error != nil {
		log.Printf("error occurred with database while checking plc id exists, get register types, plc id -> %s, drier id -> %s", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isPlcIdExists {
		return fmt.Errorf("plc id not exists"), 1, nil
	}

	isDrierIdExists, error := u.DataBaseService.CheckDrierIdExists(drierId)

	if error != nil {
		log.Printf("error occurred with database while checking direr id exists, get register types, plc id -> %s, drier id -> %s", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isDrierIdExists {
		return fmt.Errorf("drier id not exists"), 1, nil
	}

	regTypes, error := u.DataBaseService.GetAllRegisterTypes()

	if error != nil {
		log.Printf("error occurred with database while getting all register types, get register types, plc id -> %s, drier id -> %s", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	regTypesFromPlc, error := u.DataBaseService.GetRegisterTypesFromPlcByDrierId(plcId, drierId)

	if error != nil {
		log.Printf("error occurred with database while getting  register types from plc, get register types, plc id -> %s, drier id -> %s", plcId, drierId)
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	var repeatedRegisterIndex = -1

	for _, regTypeFromPlc := range regTypesFromPlc {
		repeatedRegisterIndex = -1
		for index, regType := range regTypes {
			if regTypeFromPlc == regType.Type {
				repeatedRegisterIndex = index
				break
			}
		}

		if repeatedRegisterIndex != -1 {
			regTypes = append(regTypes[:repeatedRegisterIndex], regTypes[repeatedRegisterIndex+1:]...)
		}
	}

	return nil, 0, regTypes
}
