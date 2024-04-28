package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	router := NewRouter()
	go router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))

	// Graceful Shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<- shutdown
	fmt.Println("\nServer is shutting down...")

	context, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err := router.Shutdown(context)
	if err != nil {
		router.Logger.Fatal(err)
	}
}

