package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := InitRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
