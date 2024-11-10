package admin

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/application/usecase/validation"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateAdminUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitCreateAdminUseCase(repo repository.DataBaseRepository) *CreateAdminUseCase {
	service := service.NewDataBaseService(repo)
	return &CreateAdminUseCase{
		DataBaseService: service,
	}
}

func (u *CreateAdminUseCase) Execute(request *request.Admin) (error, int) {

	if request.Email == "" {
		return fmt.Errorf("email cannot be empty"), 1
	}

	if request.Password == "" {
		return fmt.Errorf("password cannot be empty"), 1
	}

	error := validation.ValidateEmail(request.Email)

	if error != nil {
		return error, 1
	}

	error = validation.ValidatePassword(request.Password)

	if error != nil {
		return error, 1
	}

	isAdminEmailExists, error := u.DataBaseService.CheckAdminEmailExists(request.Email)

	if error != nil {
		log.Printf("error occurred with database while checking %s email exists in admin table\n", request.Email)
		return fmt.Errorf("error occurred with database"), 2
	}

	if isAdminEmailExists {
		log.Printf("%s was already exists\n", request.Email)
		return fmt.Errorf("%s was already exists", request.Email), 1
	}

	adminId := uuid.New().String()
	hashedPasswordBytes, error := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	if error != nil {
		log.Printf("failed to generate hashed password for admin %s", request.Email)
		return fmt.Errorf("failed to generate hashed password"), 2
	}

	admin := &entity.Admin{
		AdminId:  adminId,
		Email:    request.Email,
		Password: string(hashedPasswordBytes),
	}

	error = u.DataBaseService.CreateAdmin(admin)

	if error != nil {
		log.Printf("failed to create admin %s", request.Email)
		return fmt.Errorf("failed to create admin"), 2
	}

	return nil, 0
}
