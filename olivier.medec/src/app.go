package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	Connect(&m)
	r := mux.NewRouter()
	r.HandleFunc("/auth/login", Login).Methods("POST")
	r.HandleFunc("/auth/register", Register).Methods("POST")
	fmt.Println("Server running on 127.0.0.1:3000")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
