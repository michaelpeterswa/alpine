package logging

import (
	"log"

	"go.uber.org/zap"
)

func InitZap() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("unable to acquire zap logger")
	}
	return logger
}
