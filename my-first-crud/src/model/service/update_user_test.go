package service

import (
	"testing"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	mock "github.com/arturbaccarin/go-my-first-crud/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_UpdateUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when update a user returns success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "name", 98)
		userDomain.SetId(id)

		repository.EXPECT().UpdateUser(userDomain.GetId(), userDomain).Return(nil)

		err := service.UpdateUser(userDomain.GetId(), userDomain)

		assert.Nil(t, err)
	})

	t.Run("when update a user returns error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "name", 98)
		userDomain.SetId(id)

		repository.EXPECT().UpdateUser(userDomain.GetId(), userDomain).Return(rest_err.NewInternalServerError("error trying to update user"))

		err := service.UpdateUser(userDomain.GetId(), userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
	})
}
