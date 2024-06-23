package app

import (
	"applicationDesignTest/internal/controller"
	"applicationDesignTest/internal/server"
)

type controllers struct {
	CreateOrder controller.CreateOrder
}

func makeControllers(usecases server.Usecases) controllers {
	return controllers{
		CreateOrder: controller.NewCreateOrder(usecases.CreateOrder),
	}
}
