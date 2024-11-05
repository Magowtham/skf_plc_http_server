package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
)

type InitializeDataBaseUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitInitializeDataBaseUseCase(repo repository.DataBaseRepository) InitializeDataBaseUseCase {
	service := service.NewDataBaseService(repo)
	return InitializeDataBaseUseCase{
		DataBaseService: service,
	}
}

func (u *InitializeDataBaseUseCase) Execute() (error, int) {
	error := u.DataBaseService.InitializeDataBase()

	if error != nil {
		log.Printf("error occurred with database while initializing database\n")
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
