package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"powerfactors/assignment/internal/api"
)

type Server struct {
	router           *mux.Router
	logger           *zap.Logger
	httpServer       *http.Server
	timestampHandler api.TimestampHandlerInt
}

type ServerInt interface {
	Route()
	Start()
}

func NewServer(rt *mux.Router, log *zap.Logger, server *http.Server, tmHandler api.TimestampHandlerInt) Server {
	return Server{
		router:           rt,
		logger:           log,
		httpServer:       server,
		timestampHandler: tmHandler,
	}
}

func (s *Server) Route() {
	s.router.HandleFunc("/ptlist", s.timestampHandler.GetTimestamp).Methods(http.MethodGet)

	s.httpServer.Handler = s.router
}

func (s *Server) Start() {
	s.logger.Info(fmt.Sprintf("Start Server on %s\n", s.httpServer.Addr))

	err := s.httpServer.ListenAndServe()
	if err != nil {
		s.logger.Fatal("Error returned from http.listenAndServe", zap.Error(err))
		return
	}

}
