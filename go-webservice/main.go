package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GenRoute(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn(w, r)
		t := time.Since(start)
		log.Printf("%v - %v - %v\n", r.Method, r.RequestURI, t)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Requesting home")
	fmt.Fprintf(w, "hello world")
}

func main() {
	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/api/v1", GenRoute(HomeHandler))
	//	r.HandleFunc("/api/v1/health", HealthHandler)
	http.Handle("/api/v1", r)

	// Static files
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Start server
	port := "9999"
	log.Printf("Serving site on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
