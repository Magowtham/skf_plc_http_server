package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
)

type CreateRegTypeUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitCreateRegTypeUseCase(repo repository.DataBaseRepository) CreateRegTypeUseCase {
	service := service.NewDataBaseService(repo)

	return CreateRegTypeUseCase{
		DataBaseService: service,
	}
}

func (h *CreateRegTypeUseCase) Execute(regTypeRequest *request.RegisterType) (error, int) {
	if regTypeRequest.Type == "" {
		return fmt.Errorf("register type cannot be empty"), 1
	}

	if regTypeRequest.Label == "" {
		return fmt.Errorf("register label cannot be empty"), 1
	}

	isRegTypeNameExists, error := h.DataBaseService.CheckRegTypeNameExistsInRegTypes(regTypeRequest.Type)

	if error != nil {
		log.Printf("error occurred with database while checking reg type name exists, create reg type, reg type -> %s reg label -> %s\n", regTypeRequest.Type, regTypeRequest.Label)
		return fmt.Errorf("error occurred with database"), 2
	}

	if isRegTypeNameExists {
		return fmt.Errorf("register type exists"), 1
	}

	regType := entity.RegisterType{
		Type:  regTypeRequest.Type,
		Label: regTypeRequest.Label,
	}

	if error := h.DataBaseService.CreateRegType(&regType); error != nil {
		log.Printf("error occurred with database while creating the reg type, create reg type, reg type -> %s, reg label -> %s\n", regTypeRequest.Type, regTypeRequest.Label)
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
