package main

import (
	"log"

	"github.com/timwmillard/proglog/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
