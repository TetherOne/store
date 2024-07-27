package main

import (
	"log"
	"store"
)

func main() {
	srv := new(store.Server)
	if err := srv.Run("8080", nil); err != nil {
		log.Fatal(err)
	}
}
