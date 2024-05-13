package app

import (
	"applicationDesignTest/internal/config"
	"applicationDesignTest/internal/controller"
	"applicationDesignTest/internal/log"
	"applicationDesignTest/internal/model"
	"applicationDesignTest/internal/repository"
	"applicationDesignTest/internal/server"
	createorder "applicationDesignTest/internal/usecases/create_order"
	"net/http"
	"time"
)

func Start(cfg config.Config, logger log.Logger) {
	repositories := makeRepositories()
	loadFixtures(repositories.AvailabilityRepository)
	usecases := makeUsecases(repositories, logger)
	controllers := makeControllers(usecases)
	mux := createRouter(controllers)
	server := server.New(usecases, logger)
	server.Start(cfg.Server.Hostname, cfg.Server.Port, mux)
	closeRepositories(repositories)
}

type repositories struct {
	OrderRepository        repository.OrderRepository
	AvailabilityRepository repository.AvailabilityRepository
}

func makeRepositories() repositories {
	return repositories{
		OrderRepository:        repository.NewOrderRepository(),
		AvailabilityRepository: repository.NewAvailabilityRepository(),
	}
}

func loadFixtures(availabilityRepository repository.AvailabilityRepository) {
	date := func(year, month, day int) time.Time {
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}
	availabilityRecords := []model.RoomAvailability{
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 1), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 2), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 3), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 4), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 5), Quota: 0},
	}
	for i := 0; i < len(availabilityRecords); i++ {
		availabilityRepository.Push(availabilityRecords[i])
	}
}

func makeUsecases(repositories repositories, logger log.Logger) server.Usecases {
	return server.Usecases{
		CreateOrder: createorder.New(
			repositories.OrderRepository,
			repositories.AvailabilityRepository,
			logger,
		),
	}
}

type controllers struct {
	CreateOrder controller.CreateOrder
}

func makeControllers(usecases server.Usecases) controllers {
	return controllers{
		CreateOrder: controller.NewCreateOrder(usecases.CreateOrder),
	}
}

func createRouter(controllers controllers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/order", controllers.CreateOrder.Handle)
	return mux
}

func closeRepositories(repositories repositories) {
	repositories.AvailabilityRepository.Close()
	repositories.OrderRepository.Close()
}
