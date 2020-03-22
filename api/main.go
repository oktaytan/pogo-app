package main

import (
	"log"
	"net/http"
	"pogo/api/routers"

	"github.com/urfave/negroni"
)

func main() {

	// Router initialize
	router := routers.InitRoutes()

	// Handle negroni
	n := negroni.Classic()
	n.UseHandler(router)

	log.Println("Server started on port 5000")

	// Server listening
	http.ListenAndServe(":5000", n)
}
