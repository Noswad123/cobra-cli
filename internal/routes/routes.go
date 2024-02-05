package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/api/data", apiDataHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some data from the API"
	fmt.Fprintln(w, data)
}