package service

import (
	"testing"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"github.com/arturbaccarin/go-my-first-crud/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when user already exists returns error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "name", 98)
		userDomain.SetId(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUser(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "email is already registered in another account")
	})

	t.Run("when user is not registered returns error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "name", 98)
		userDomain.SetId(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewInternalServerError("error trying to create user"))

		user, err := service.CreateUser(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create user")
	})

	t.Run("when user is not registered returns success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "name", 98)
		userDomain.SetId(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateUser(userDomain)

		assert.NotNil(t, user)
		assert.Nil(t, err)
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
	})
}
