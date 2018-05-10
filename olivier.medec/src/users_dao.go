package main

import (
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
func (m *UsersDAO) FindByID(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
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
