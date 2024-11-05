package user

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

type UserLoginUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitUserLoginUseCase(dbRepo repository.DataBaseRepository) UserLoginUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return UserLoginUseCase{
		DataBaseService: dbService,
	}
}

func (u *UserLoginUseCase) Execute(userRequest *request.UserLogin) (error, int, string) {
	if userRequest.Email == "" {
		return fmt.Errorf("user email cannot be empty"), 1, ""
	}

	if userRequest.Password == "" {
		return fmt.Errorf("user password cannot be empty"), 1, ""
	}

	if error := validation.ValidateEmail(userRequest.Email); error != nil {
		return error, 1, ""
	}

	isUserEmailExists, error := u.DataBaseService.CheckUserEmailExists(userRequest.Email)

	if error != nil {
		log.Printf("error occurred with database while checking user email exists, user login, user email -> %s", userRequest.Email)
		return fmt.Errorf("error occurred with database"), 2, ""
	}

	if !isUserEmailExists {
		return fmt.Errorf("user email not exists"), 1, ""
	}

	user, error := u.DataBaseService.GetUserByEmail(userRequest.Email)

	if error != nil {
		log.Printf("error occurred with database while getting user by email, user log in, user email -> %s", userRequest.Email)
		return fmt.Errorf("error occurred with database"), 2, ""
	}

	if error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); error != nil {
		return fmt.Errorf("incorrect password"), 1, ""
	}

	secreteKey := os.Getenv("SECRETE_KEY")

	if secreteKey == "" {
		log.Printf("missing env variable SECRETE_KEY")
		return fmt.Errorf("error occurred while generating the token"), 2, ""
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserId,
		"email":   user.Email,
		"label":   user.Label,
		"exp":     time.Now().Add(time.Hour * 24 * 360).Unix(),
	})

	tokenString, error := token.SignedString([]byte(secreteKey))

	if error != nil {
		log.Printf("error occurred while generating jwt token")
		return fmt.Errorf("error occurred while generating token"), 2, ""
	}

	return nil, 0, tokenString
}
