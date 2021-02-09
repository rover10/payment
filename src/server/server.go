package server

import (
	"net/http"

	"github.com/rover10/payment/config"
	"github.com/rover10/payment/database"
)

type Server struct {
	Config   config.Config
	Database database.DBClient
}

func (s *Server) Start() error {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Errorf("Error start server: %v", err)
	}
	return err
}
