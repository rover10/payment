package main

import (
	"fmt"

	"github.com/rover10/payment/lib/api"
	"github.com/rover10/payment/lib/config"
	"github.com/rover10/payment/lib/database"
)

func main() {
	fmt.Println("Payment service")
	cfg := config.Config{Host: "localhost", Port: 8080}
	initDB(&cfg)
	api.StartServer(&cfg)
}
func initDB(cfg *config.Config) database.DBClient {
	//Connect to db
	database.DB = &database.Client{}
	database.DB.DBConnect(cfg)
	return database.DB
	//server.Database = &dbclient

}
