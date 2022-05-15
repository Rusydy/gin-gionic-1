package users

import (
	"errors"
	"gin-gionic-1/connection"
	user "gin-gionic-1/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// UserCollection statically declared
const UserCollection = "user"

var (
	errNotExist        = errors.New("Users are not exist")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidBody     = errors.New("Invalid request body")
	errInsertionFailed = errors.New("Error in the user insertion")
	errUpdationFailed  = errors.New("Error in the user updation")
	errDeletionFailed  = errors.New("Error in the user deletion")
)

// GetAllUsers returns all users
func GetAllUser(c *gin.Context) {
	// Get DB from Mongo Config
	db := connection.GetMongoDB()
	users := user.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "users": &users})
}
