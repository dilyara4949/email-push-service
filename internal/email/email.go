package email

import (
	"fmt"
	"github.com/dilyara4949/email-push-service/internal/config"
	"log/slog"
	"net/smtp"
)

type EmailService struct {
	cfg    config.Config
	logger *slog.Logger
}

func NewEmailService(cfg config.Config, logger *slog.Logger) *EmailService {
	return &EmailService{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *EmailService) SendEmail(emailID string, message string) error {
	from := s.cfg.Email.Username
	to := emailID
	auth := smtp.PlainAuth("", s.cfg.Email.Username, s.cfg.Email.Password, s.cfg.Email.SMTPHost)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", s.cfg.Email.SMTPHost, s.cfg.Email.SMTPPort),
		auth,
		from,
		[]string{to},
		[]byte(message),
	)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	s.logger.Info("email sent successfully", "to", emailID)
	return nil
}
