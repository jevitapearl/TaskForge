package main

import (
	"log"
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/server"
)

func main() {
	s := server.New()

	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", s.Router())
	if err != nil {
		panic(err)
	}

}
