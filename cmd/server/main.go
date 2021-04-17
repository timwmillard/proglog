package main

import (
	"log"

	"github.com/timwmillard/proglog/server"
)

const bindAddress = ":8080"

func main() {
	srv := server.NewHTTPServer(bindAddress)
	log.Printf("Starting proglog server at %v", bindAddress)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
