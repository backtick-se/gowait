package util

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var LogModule = fx.Module(
	"log",
	fx.Provide(NewLogger),
	fx.Provide(NewSugaredLogger),
)

func NewLogger(lc fx.Lifecycle) *zap.Logger {
	logger, _ := zap.NewDevelopment()

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			// flushes buffer, if any
			logger.Sync()
			return nil
		},
	})

	return logger
}

func NewSugaredLogger(logger *zap.Logger) *zap.SugaredLogger {
	return logger.Sugar()
}
