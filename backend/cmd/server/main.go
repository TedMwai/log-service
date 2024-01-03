package main

import "log-management/config"


func main() {
	config.Init()

	service := New();

	// Starts the service in a go routine
	service.Start()

	// Shutdown blocks until a shutdown signal is received
	service.Shutdown()
}

