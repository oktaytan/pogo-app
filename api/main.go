package main

import (
	"net/http"
	"pogo/api/routers"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})

	// Router initialize
	router := routers.InitRoutes()

	// Handle negroni
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)

	// Server listening
	http.ListenAndServe(":5000", n)
}
