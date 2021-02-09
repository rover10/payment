package main

import (
	"fmt"

	"github.com/rover10/payment/lib/api"
)

func main() {
	fmt.Println("Payment service")
	api.StartServer()
}
