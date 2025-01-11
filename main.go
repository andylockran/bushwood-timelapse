package main

import (
	"log"
	"net/http"

	"github.com/andylockran/bushwood-timelapse/handlers"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server.  Listening at http://localhost%s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
