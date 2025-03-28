package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("%+v\n", r)
	// fmt.Printf("Method: %v\n", r.Method)
	// fmt.Printf("Url: %v\n", r.URL)
	respondWithJSON(w, 200, struct{}{})
}
