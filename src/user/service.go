package user

import (
	"blog/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	Collection *mongo.Collection
}

// Constructor
func InitStateService() *UserService {
	var stateService UserService
	stateService.Collection = utils.Db().Collection("users")

	return &stateService
}

func (service *UserService) GetUserName(userName string) (*User, error) {
	var state User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := service.Collection.FindOne(ctx, bson.M{"chatid": userName}).Decode(&state)

	return &state, err
}

func (service *UserService) Upsert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	opts := options.Replace().SetUpsert(true)
	_, err := service.Collection.ReplaceOne(ctx, bson.D{{"userName", user.UserName}}, user, opts)

	return err
}
