package logger

import (
	"context"
	"pulse/utils/config"
	_const "pulse/utils/const"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger(lc fx.Lifecycle, env *config.Env) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	if env.DeploymentEnv == _const.Deployment_Production {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})
	return logger, nil
}
