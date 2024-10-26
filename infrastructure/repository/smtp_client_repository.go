package repository

import (
	"net/smtp"

	"github.com/VsenseTechnologies/skf_plc_http_server/infrastructure/smtpclient"
)

type SmtpClientRepository struct {
	smtpClient *smtpclient.SmtpClient
}

func NewSmtpClientRepository(smtpClient *smtpclient.SmtpClient) SmtpClientRepository {
	return SmtpClientRepository{
		smtpClient,
	}
}

func (repo *SmtpClientRepository) SendEmail(emails []string, body []byte) error {
	error := smtp.SendMail(repo.smtpClient.SmtpServerAddress, repo.smtpClient.SmtpAuth, repo.smtpClient.SmtpUserName, emails, body)
	return error
}
