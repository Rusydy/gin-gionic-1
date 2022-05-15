package models

import (
	"gin-gionic-1/connection"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// user struct
type User struct {
	ID        bson.ObjectId `bson:"_id"`
	FirstName string        `bson:"first_name"`
	LastName  string        `bson:"last_name"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

// users list
type Users []User

// UserInfo model function
func UserInfo(id bson.ObjectId, userCollection string) (User, error) {
	// Get DB from Mongo Config
	db := connection.GetMongoDB()
	user := User{}
	err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)
	return user, err
}
