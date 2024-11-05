package admin

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/service"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
	"golang.org/x/crypto/bcrypt"
)

type GiveUserAccessUseCase struct {
	DataBaseService   *service.DataBaseService
	SmtpClientService service.SmtpClientService
}

func InitGiveUserAccessUseCase(dbRepo repository.DataBaseRepository, smtpRepo repository.SmtpClientRepository) GiveUserAccessUseCase {
	dataBaseService := service.NewDataBaseService(dbRepo)
	smtpClientService := service.NewSmtpClientService(smtpRepo)
	return GiveUserAccessUseCase{
		DataBaseService:   dataBaseService,
		SmtpClientService: smtpClientService,
	}
}

func (u *GiveUserAccessUseCase) Execute(userAccessRequest *request.UserAccess) (error, int) {
	if userAccessRequest.UserId == "" {
		return fmt.Errorf("user id cannot be empty"), 1
	}

	if userAccessRequest.Password == "" {
		return fmt.Errorf("password cannot be empty"), 1
	}

	isUserIdExists, error := u.DataBaseService.CheckUserIdExists(userAccessRequest.UserId)

	if error != nil {
		log.Printf("error occurred with database while checking user id exists, give user access, user id -> %s", userAccessRequest.UserId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if !isUserIdExists {
		return fmt.Errorf("user id not exists"), 1
	}

	user, error := u.DataBaseService.GetUserById(userAccessRequest.UserId)

	if error != nil {
		log.Printf("error occurred with database while getting the user by id, give user access, user id -> %s", user.UserId)
		return fmt.Errorf("error occurred with database"), 2
	}

	if error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userAccessRequest.Password)); error != nil {
		return fmt.Errorf("incorrect password"), 1
	}

	emailTemplate, error := template.ParseFiles("template/email_template.html")

	if error != nil {
		log.Printf("error occurred while parsing the email html template file,give user access, user email -> %s user id -> %s", user.Email, userAccessRequest.UserId)
		return fmt.Errorf("error occurred while parsing"), 2
	}

	var emailBody bytes.Buffer

	emailTemplate.Execute(&emailBody, struct {
		Email    string
		Password string
	}{
		Email:    user.Email,
		Password: userAccessRequest.Password,
	})

	emailSubject := "Access to SKF Elixer Application"

	emailMessage := []byte(fmt.Sprintf("Subject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s", emailSubject, emailBody.String()))

	if error := u.SmtpClientService.SendEmail([]string{user.Email}, []byte(emailMessage)); error != nil {
		log.Printf("error occurred while sending the email, give use access, user email -> %s user id -> %s", user.Email, userAccessRequest.UserId)
		return fmt.Errorf("error occurred while send email"), 2
	}

	return nil, 0
}
