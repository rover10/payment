package server

import (
	"config"
	"database"
)

type Server struct {
	Config   config.Config
	Database database.DBClient
}
