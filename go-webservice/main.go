package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Requesting home")
	fmt.Fprintf(w, "hello world we are a go app")
}

func main() {
	// API Routes
	r := mux.NewRouter()
	r.HandleFunc("/api/", HomeHandler)

	// Where React is served from
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	http.Handle("/", r)

	// Middlewares
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	// Start server
	port := "9999"
	log.Printf("Serving site on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, loggedRouter))
}
