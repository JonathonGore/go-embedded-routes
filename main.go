package main

import (
	"log"
	"net/http"

	"github.com/JonathonGore/go-embedded-routes/handlers"
	"github.com/JonathonGore/go-embedded-routes/server"
)

func main() {
	api, err := handlers.New()
	if err != nil {
		log.Fatalf("unable to create handler: %v", err)
	}

	s, err := server.New(api)
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}

	srv := &http.Server{
		Addr:      ":3000",
		Handler:   s,
	}

	log.Fatal(srv.ListenAndServe())
}
