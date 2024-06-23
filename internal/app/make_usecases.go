package app

import (
	"applicationDesignTest/internal/log"
	"applicationDesignTest/internal/server"
	createorder "applicationDesignTest/internal/usecase/create_order"
)

func makeUsecases(repositories repositories, logger log.Logger) server.Usecases {
	return server.Usecases{
		CreateOrder: createorder.New(
			repositories.OrderRepository,
			repositories.AvailabilityRepository,
			logger,
		),
	}
}
