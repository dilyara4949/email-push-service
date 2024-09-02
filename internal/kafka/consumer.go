package kafka

import (
	"context"
	"github.com/dilyara4949/email-push-service/internal/config"
	"github.com/dilyara4949/email-push-service/internal/email"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

type Consumer struct {
	reader       *kafka.Reader
	emailService *email.EmailService
	logger       *slog.Logger
}

func NewConsumer(cfg config.Config, emailService *email.EmailService, logger *slog.Logger) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Kafka.Brokers},
		Topic:   cfg.Kafka.Topic,
		GroupID: cfg.Kafka.GroupID,
	})

	return &Consumer{
		reader:       r,
		emailService: emailService,
		logger:       logger,
	}
}

func (c *Consumer) Start(ctx context.Context) {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			c.logger.Info("Error reading message:", err)
			continue
		}

		emailID := string(m.Key)
		message := string(m.Value)
		err = c.emailService.SendEmail(emailID, message)
		if err != nil {
			c.logger.Error("failed to send email", "error", err)
		}
	}
}
