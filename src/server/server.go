package server

import (
	"errors"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/rover10/payment/src/config"
	"github.com/rover10/payment/src/database"
)

type Server struct {
	Config   config.Config
	Database database.DBClient
}

func (s *Server) Start() error {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Errorf("Error start server: %v", err)
	}
	return errors.New("Server failed")
}
