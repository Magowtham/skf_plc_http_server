package smtpclient

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type SmtpClient struct {
	SmtpUserName      string
	SmtpServerAddress string
	SmtpAuth          smtp.Auth
}

func SetupClient() SmtpClient {
	smtpUserName := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpServiceHost := os.Getenv("SMTP_SERVICE_HOST")
	smtpServicePort := os.Getenv("SMTP_SERVICE_PORT")

	if smtpUserName == "" {
		log.Fatalln("missing the env variable SMTP_USERNAME")
	}

	if smtpPassword == "" {
		log.Fatalln("missing the env variable SMTP_PASSWORD")
	}

	if smtpServiceHost == "" {
		log.Fatalln("missing the env variable SMTP_SERVICE_HOST")
	}

	if smtpServicePort == "" {
		log.Fatalln("missing the env variable SMTP_SERVICE_PORT")
	}

	smtpAuth := smtp.PlainAuth("", smtpUserName, smtpPassword, smtpServiceHost)

	return SmtpClient{
		SmtpUserName:      smtpUserName,
		SmtpServerAddress: fmt.Sprintf("%s:%s", smtpServiceHost, smtpServicePort),
		SmtpAuth:          smtpAuth,
	}
}
