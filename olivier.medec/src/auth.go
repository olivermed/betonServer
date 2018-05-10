package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
)

//DAO du user
var m UsersDAO

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func GetNewToken(key string) string {
	keyToken := []byte("hmacSampleSecret" + key)
	date := time.Now().AddDate(1, 0, 0).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": date,
	})

	tokenString, err := token.SignedString(keyToken)

	fmt.Println(tokenString, err)
	return tokenString
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userAuth User

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := json.Unmarshal(b, &userAuth); err != nil {
		fmt.Println(userAuth)
		respondWithError(w, http.StatusBadRequest, "Invalid user ID 1")
		return
	}

	user, err := m.FindUser(userAuth.Email, userAuth.Password)
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
	user.Token = GetNewToken(user.Email)
	if err := m.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}
