package tests

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.9.0/mongo

// Fixture -- is a fabric function.
// Фабричная функция -- это (creational pattern) паттерн,
// который возвращает новый экземпляр какого-то объекта.

func ClearDb(collections []*mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for _, collection := range collections {
		err := collection.Drop(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func newUser(userName string, userCollection *mongo.Collection) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := userCollection.InsertOne(ctx, bson.D{
		{"userName", userName},
		{"firstName", "dada"},
		{"lastName", "yaya"},
	})
	return res, err
}

func SetupStateCollection(userCollection *mongo.Collection) []*mongo.InsertOneResult {

	state1, err := newUser("Vlad", userCollection)
	if err != nil {
		log.Fatal(err)
	}

	state2, err := newUser("Denis", userCollection)
	if err != nil {
		log.Fatal(err)
	}

	state3, err := newUser("Pens", userCollection)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(state3.InsertedID)

	// Example how to obtain state id:
	// id := res.InsertedID
	return []*mongo.InsertOneResult{state1, state2, state3}
}

func InitApp() *gin.Engine {
	// Gin Init
	app := gin.New()
	app.Use(gin.Recovery())

	return app
}
