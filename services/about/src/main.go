package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var About = "By Tony Bailey"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "healthy")
	}).Methods("GET")

	r.HandleFunc("/health/readiness", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "healthy")
	}).Methods("GET")

	r.HandleFunc("/health/liveness", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "healthy")
	}).Methods("GET")

	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, About)
	}).Methods("GET")

	http.ListenAndServe(":5000", r)
}
