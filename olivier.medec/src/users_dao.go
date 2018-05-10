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
func (m *UsersDAO) FindUser(email string, password string) (User, error) {
	var user User
	fmt.Println("Connection user :: ", email, password)
	if errUpdate := db.C(COLLECTION).UpdateId(user.ID, bson.M{"token": GetNewToken(email)}); errUpdate == nil {
		return user, errUpdate
	}
	err := db.C(COLLECTION).Find(bson.M{"email": email, "password": password}).One(&user)

	return user, err
}

//Insert push user into db
func (m UsersDAO) Insert(user User) error {
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
