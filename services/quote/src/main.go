package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Quote = "Optimism is an occupational hazard of programming: feedback is the treatment."

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, Quote)
	}).Methods("GET")

	http.ListenAndServe(":5000", r)
}
