package main

import (
	"log"
	"net/http"
	"pogo/api/routers"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
	})

	// Router initialize
	router := routers.InitRoutes()

	// Handle negroni
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)

	log.Println("Server started on port 5000")

	// Server listening
	http.ListenAndServe(":5000", n)
}
