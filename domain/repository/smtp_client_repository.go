package repository

type SmtpClientRepository interface {
	SendEmail(emails []string, body []byte) error
}
