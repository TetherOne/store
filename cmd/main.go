package main

import (
	"log"
	"store/todo"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatal(err)
	}
}
