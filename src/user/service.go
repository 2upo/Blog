package user

import (
	"blog/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Collection *mongo.Collection
}

// Constructor
func InitUserService() *UserService {
	var userService UserService
	userService.Collection = utils.Db().Collection("users")

	return &userService
}

func (service *UserService) GetUserName(userName string) (*User, error) {
	var name User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := service.Collection.FindOne(ctx, bson.M{"userName": userName}).Decode(&name)

	return &name, err
}

func (service *UserService) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := service.Collection.InsertOne(ctx, user)

	return err

}
