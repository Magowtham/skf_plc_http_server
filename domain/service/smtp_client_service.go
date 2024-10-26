package service

import "github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"

type SmtpClientService struct {
	smtpRepo repository.SmtpClientRepository
}

func NewSmtpClientService(repo repository.SmtpClientRepository) SmtpClientService {
	return SmtpClientService{
		smtpRepo: repo,
	}
}

func (service *SmtpClientService) SendEmail(emails []string, body []byte) error {
	return service.smtpRepo.SendEmail(emails, body)
}
