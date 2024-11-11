package user

import (
	"fmt"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
)

type CreateUserFeedbackUseCase struct {
	DataBaseService *service.DataBaseService
}

func InitCreateUserFeedbackUseCase(dbRepo repository.DataBaseRepository) *CreateUserFeedbackUseCase {
	dbService := service.NewDataBaseService(dbRepo)
	return &CreateUserFeedbackUseCase{
		DataBaseService: dbService,
	}
}

func (u *CreateUserFeedbackUseCase) Execute(userId string, userRequest *request.UserFeedback) (error, int) {
	if userId == "" {
		return fmt.Errorf("user id cannot be empty"), 1
	}

	if userRequest.Feedback == "" {
		return fmt.Errorf("feedback cannot be empty"), 1
	}

	isUserIdExists, err := u.DataBaseService.CheckUserIdExists(userId)

	if err != nil {
		log.Printf("error occurred with database while checking user id exists, user feedback, user id -> %v, error -> %v\n", userId, err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isUserIdExists {
		return fmt.Errorf("user id not exists"), 1
	}

	if err := u.DataBaseService.CreateUserFeedback(userId, userRequest.Feedback); err != nil {
		log.Printf("error occurred with database while creating the user feedback, user feedback, user id -> %v, error -> %v\n", userId, err.Error())
		return fmt.Errorf("error occurred with database"), 2
	}

	return nil, 0
}
