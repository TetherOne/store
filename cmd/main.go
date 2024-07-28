package main

import (
	"log"
	"store"
	"store/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(store.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
