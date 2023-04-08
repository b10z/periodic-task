package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"powerfactors/assignment/internal/api"
)

type Listener struct {
	router           *mux.Router
	logger           *zap.Logger
	httpServer       *http.Server
	timestampHandler api.TimestampHandlerInt
}

type ListenerInt interface {
	Route()
	Start()
}

func NewListener(rt *mux.Router, log *zap.Logger, server *http.Server, tmHandler api.TimestampHandlerInt) Listener {
	return Listener{
		router:           rt,
		logger:           log,
		httpServer:       server,
		timestampHandler: tmHandler,
	}
}

func (s *Listener) Route() {
	s.router.HandleFunc("/ptlist", s.timestampHandler.GetTimestamp).Methods(http.MethodGet)

	s.httpServer.Handler = s.router
}

func (s *Listener) Start() {
	s.logger.Info(fmt.Sprintf("Start Listener on %s\n", s.httpServer.Addr))

	err := s.httpServer.ListenAndServe()
	if err != nil {
		s.logger.Fatal("Error returned from http.listenAndServe", zap.Error(err))
		return
	}
}
