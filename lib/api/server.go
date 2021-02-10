package api

import (
	"fmt"
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/rover10/payment/lib/config"
	"github.com/rover10/payment/lib/database"
)

type Server struct {
	Config   *config.Config
	Database *database.Client
	Router   *echo.Echo
	Host     string
	Port     int
}

func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)
	log.Infof("Listening on %s", address)
	return s.Router.Start(address)
}

func NewServer(cfg *config.Config) *Server {
	server := Server{Config: cfg, Router: echo.New()}
	server.Router.GET(path.Join(cfg.APIPath, "/v1/payment/history/:userId"), server.TransactionHistory)
	return &server
}

func StartServer(cfg *config.Config) {
	server := NewServer(cfg)
	server.Config = cfg
	server.Start()
}
