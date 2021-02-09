package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/rover10/payment/lib/config"
	"github.com/rover10/payment/lib/database"
)

type Server struct {
	Config   config.Config
	Database database.DBClient
	router   *echo.Echo
}

func (s *Server) Start() error {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Errorf("Error start server: %v", err)
	}
	return errors.New("Server failed")
}
