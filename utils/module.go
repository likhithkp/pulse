package utils

import (
	"pulse/utils/config"
	"pulse/utils/jwt"
	"pulse/utils/logger"
	"pulse/utils/middleware"
	"pulse/utils/other"
	"pulse/utils/server"

	"go.uber.org/fx"
)

var Module = fx.Module("utils",
	config.Module,
	jwt.Module,
	logger.Module,
	middleware.Module,
	other.Module,
	server.Module,
)
