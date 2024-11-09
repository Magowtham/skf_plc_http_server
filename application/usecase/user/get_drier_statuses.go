package user

import (
	"fmt"
	"log"
	"strings"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type GetDrierStatusesUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitGetDrierStatusesUseCase(dbRepo repository.DataBaseRepository) *GetDrierStatusesUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &GetDrierStatusesUseCase{
		DataBaseService: dbService,
	}
}

func (u *GetDrierStatusesUseCase) Execute(plcId string, drierId string) (error, int, []*entity.DrierStatus) {
	if plcId == "" {
		return fmt.Errorf("plc id cannot be empty"), 1, nil
	}

	if drierId == "" {
		return fmt.Errorf("drier id cannot be empty"), 1, nil
	}

	isPlcIdExists, err := u.DataBaseService.CheckPlcIdExists(plcId)

	if err != nil {
		log.Printf("error occurred with database while checking plc id exists, get drier statuses, plc id -> %v, drier id -> %v, Error -> %v\n", plcId, drierId, err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isPlcIdExists {
		return fmt.Errorf("plc id not exists"), 1, nil
	}

	isDrierIdExists, err := u.DataBaseService.CheckDrierIdExists(drierId)

	if err != nil {
		log.Printf("error occurred with database while checking drier id exists, get drier statuses, plc id -> %v, drier id -> %v, Error -> %v\n", plcId, drierId, err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	if !isDrierIdExists {
		return fmt.Errorf("drier id not exists"), 1, nil
	}
	registerTypes, err := u.DataBaseService.GetAllRegisterTypes()

	if err != nil {
		log.Printf("error occurred with database while getting all the register types, plc id -> %v drier id -> %v, Error -> %v\n", plcId, drierId, err.Error())
		return fmt.Errorf("error occurred with database"), 2, nil
	}

	var drierStatuses []*entity.DrierStatus

	for _, regType := range registerTypes {
		if strings.Split(regType.Type, "_")[0] == "st" {
			regValue, err := u.DataBaseService.GetRegisterValueByRegisterTypeAndDrierId(plcId, drierId, regType.Type)
			if err != nil {
				log.Printf("error occurred with database while getting the register value, get drier statuses, plc id -> %v, drier id -> %v, Error -> %v\n", plcId, drierId, err.Error())
				return fmt.Errorf("error occurred with database"), 2, nil
			}

			drierStatuses = append(drierStatuses, &entity.DrierStatus{
				RegisterType:  regType.Type,
				RegisterValue: regValue,
			})
		}
	}

	return nil, 0, drierStatuses
}
