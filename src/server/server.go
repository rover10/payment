package server

import (
	"config"
	"database"
	"net/http"
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
