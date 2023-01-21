package adapters

import (
	"github.com/protohedge/protohedge.api/internal/config"
	"go.uber.org/zap"
)

func NewLogger(cfg *config.Config) *zap.Logger {
	switch cfg.Env {
	case "production":
		logger, err := zap.NewProduction()

		if err != nil {
			panic(err)
		}

		return logger
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return logger
}
