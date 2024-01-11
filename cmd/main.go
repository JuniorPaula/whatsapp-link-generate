package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	port := ":8080"

	srv := &http.Server{
		Addr:              port,
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	log.Println("Server running on port", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
