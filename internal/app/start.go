package app

import (
	"applicationDesignTest/internal/config"
	"applicationDesignTest/internal/log"
	"applicationDesignTest/internal/server"
	"net/http"
)

func Start(cfg config.Config, logger log.Logger) {
	repositories := makeRepositories()
	usecases := makeUsecases(repositories, logger)
	controllers := makeControllers(usecases)
	mux := createRouter(controllers)
	server := server.New(usecases, logger)
	server.Start(cfg.Server.Hostname, cfg.Server.Port, mux)
	closeRepositories(repositories)
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
