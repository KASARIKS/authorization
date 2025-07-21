package main

import (
	"log"
	"net/http"

	"github.com/kasariks/authorization/internal/server"
)

func main() {
	logger := &log.Logger{}
	router, err := server.CreateServer(logger)
	if err != nil {
		log.Fatal(err)
	}
	if err = http.ListenAndServe(router.Server.Addr, router.Server.Handler); err != nil {
		log.Fatal(err)
	}
}
