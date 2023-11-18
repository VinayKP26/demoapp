package main

import (
	"context"
	"os"

	"github.com/VinayKP26/demoapp/router"
)

func main() {
	ctx := context.Background()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	println("Starting the server on port: ", port)
	router.StartServer(port, ctx)
}
