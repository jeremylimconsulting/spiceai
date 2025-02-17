package loggers

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
)

var (
	zapLogger *zap.Logger
)

func ZapLogger() *zap.Logger {
	if zapLogger != nil {
		return zapLogger
	}

	var err error
	if os.Getenv("SPICE_DEBUG") == "1" {
		zapLogger, err = zap.NewDevelopment()
	} else {
		zapLogger, err = zap.NewProduction()
	}
	if err != nil {
		// Fall back to standard logging
		log.Println(fmt.Errorf("unable to create Zap logger: %w", err))
		return nil
	}

	return zapLogger
}

func ZapLoggerSync() {
	if zapLogger != nil {
		err := zapLogger.Sync()
		if err != nil {
			// Swallow errors in sync
			// https://github.com/uber-go/zap/issues/880
			return
		}
	}
}
