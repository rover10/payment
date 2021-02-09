package main

import (
	"fmt"

	"github.com/rover10/payment/server"
)

func main() {
	fmt.Println("Payment service")
	server.Start()
}
