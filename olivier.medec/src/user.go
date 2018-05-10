package main

import "github.com/globalsign/mgo/bson"

type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Nom      string        `bson:"nom" json:"nom"`
	Prenom   string        `bson:"prenom" json:"prenom"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	Token    string        `bson:"token" json:"token"`
}
