package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//UsersDAO Data Access Object
type UsersDAO struct {
	Server   string
	Database string
}

const COLLECTION string = "Users"

var db *mgo.Database

//FindByID finds user by id
func (m *UsersDAO) FindUser(email string) (User, error) {
	var user User
	fmt.Println("Connection user :: ", email)
	err := db.C(COLLECTION).Find(bson.M{"email": email}).One(&user)

	return user, err
}

//Insert push user into db
func (m UsersDAO) Insert(user User) error {
	fmt.Println("New user :: ", user.Email)
	err := db.C(COLLECTION).Insert(&user)
	return err
}

//Connect to db
func Connect(m *UsersDAO) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
