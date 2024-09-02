package main

import (
	"context"
	"github.com/dilyara4949/email-push-service/internal/config"
	"github.com/dilyara4949/email-push-service/internal/email"
	"github.com/dilyara4949/email-push-service/internal/kafka"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	emailService := email.NewEmailService(cfg, logger)
	kafkaConsumer := kafka.NewConsumer(cfg, emailService, logger)

	kafkaConsumer.Start(ctx)
}
