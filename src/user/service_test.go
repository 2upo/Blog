package user

import (
	"blog/tests"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userService = InitUserService()

func TestInsert(t *testing.T) {
	ass := assert.New(t)

	new_user := User{
		UserName:  "testuserName",
		FirstName: "testuserFirst",
		LastName:  "testuserLast",
	}

	err := userService.Insert(&new_user)
	ass.Nil(err)

	user, err := mockGetByName("testuserName")

	ass.Nil(err)
	ass.Equal(user.UserName, new_user.UserName)
	ass.Equal(user.LastName, new_user.LastName)

	tests.ClearDb([]*mongo.Collection{userService.Collection})
}

func mockGetByName(name string) (*User, error) {
	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := userService.Collection.FindOne(ctx, bson.M{"username": name}).Decode(&user)

	return &user, err
}
