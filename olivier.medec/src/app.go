package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

type userID struct {
	ID string `json:"id"`
}

var m UsersDAO

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Login the user
func Login(w http.ResponseWriter, r *http.Request) {
	var userid userID

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := json.Unmarshal(b, &userid); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID 1")
		return
	}

	user, err := m.FindByID(userid.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID 2")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// Register users
func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := m.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
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
