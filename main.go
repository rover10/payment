package main

import (
	"fmt"

	"github.com/rover10/payment/src/server"
)

func main() {
	fmt.Println("Payment service")
	s := server.Server{}
	s.Start()
}
