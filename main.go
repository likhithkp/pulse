package main

import (
	"pulse/application"
	"pulse/data_access"
	"pulse/domain"
	"pulse/utils"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		application.Module,
		data_access.Module,
		domain.Module,
		utils.Module,
	).Run()
}
