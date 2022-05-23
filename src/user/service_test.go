package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestInsert(t *testing.T) {
	ass := assert.New(t)

	new_user := User{}

	err := userService.Insert(&new_user)
	ass.Nil(err)

	state, err := mockGetById(new_user.ID)

	ass.Nil(err)
	ass.Equal(state.Header, new_user.Header)
	ass.Equal(state.Content, new_user.Content)

	tests.ClearDb([]*mongo.Collection{userService.Collection})
}
