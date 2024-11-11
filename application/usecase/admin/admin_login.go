package admin

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/VsenseTechnologies/skf_plc_http_server/application/usecase/validation"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AdminLoginUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitAdminLoginUseCase(repo repository.DataBaseRepository) *AdminLoginUseCase {
	service := service.NewDataBaseService(repo)

	return &AdminLoginUseCase{
		DataBaseService: service,
	}
}

func (u *AdminLoginUseCase) Execute(adminRequest request.AdminLogin) (error, int, string) {
	if adminRequest.Email == "" {
		return fmt.Errorf("email cannot be empty"), 1, ""
	}

	if adminRequest.Password == "" {
		return fmt.Errorf("password cannot be empty"), 1, ""
	}

	error := validation.ValidateEmail(adminRequest.Email)

	if error != nil {
		return error, 1, ""
	}

	isAdminEmailExists, error := u.DataBaseService.CheckAdminEmailExists(adminRequest.Email)

	if error != nil {
		log.Printf("error occurred with database while checking admin email exists -> %s", adminRequest.Email)
		return fmt.Errorf("error occurred with database"), 2, ""
	}

	if !isAdminEmailExists {
		return fmt.Errorf("email not exists"), 1, ""
	}

	admin, error := u.DataBaseService.GetAdminByEmail(adminRequest.Email)

	if error != nil {
		log.Printf("error occurred with database while getting admin by email -> %s", adminRequest.Email)
		return fmt.Errorf("error occurred with database"), 2, ""
	}

	error = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(adminRequest.Password))

	if error != nil {
		return fmt.Errorf("incorrect password"), 1, ""
	}

	secreteKey := os.Getenv("S1_SECRETE_KEY")

	if secreteKey == "" {
		log.Printf("missing env variable S1_SECRETE_KEY")
		return fmt.Errorf("error occurred while generating the token"), 2, ""
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": admin.AdminId,
		"email":    admin.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 360).Unix(),
	})

	tokenString, error := token.SignedString([]byte(secreteKey))

	if error != nil {
		log.Printf("error occurred while generating jwt token")
		return fmt.Errorf("error occurred while generating token"), 2, ""
	}

	return nil, 0, tokenString
}
