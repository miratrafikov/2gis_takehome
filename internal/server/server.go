package server

import (
	"applicationDesignTest/internal/log"
	createorder "applicationDesignTest/internal/usecases/create_order"
	"errors"
	"fmt"
	"net/http"
)

type Server struct {
	usecases Usecases
	logger   log.Logger
}

type Usecases struct {
	CreateOrder createorder.Usecase
}

func New(usecases Usecases, logger log.Logger) Server {
	return Server{
		usecases: usecases,
		logger:   logger,
	}
}

func (s Server) Start(hostname string, port int, mux http.Handler) {
	host := fmt.Sprintf("%s:%d", hostname, port)
	s.logger.Infof("Starting server on %s...", host)
	s.handleStoppage(http.ListenAndServe(host, mux))
}

func (s Server) handleStoppage(err error) {
	if errors.Is(err, http.ErrServerClosed) {
		s.logger.Info("Server stopped")
	} else if err != nil {
		s.logger.Fatalf("Server failed: %s", err)
	}
}
