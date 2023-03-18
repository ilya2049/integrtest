package main

import (
	"integrtest/startup"

	"log"
)

func main() {
	httpServer := startup.InitializeHTTPServer()

	if err := httpServer.Listen(":3123"); err != nil {
		log.Fatal("test server listening error:", err)
	}

	// TODO: implement graceful shutdown
}
