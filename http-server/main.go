package main

import (
	"github.com/jabernardo/tuts/http-server/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", handlers.HomeRoute)
	mux.HandleFunc("/articles/{article}", handlers.ArticlesRoute)

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		log.Fatalln("Cannot run the server", err)
	}
}
