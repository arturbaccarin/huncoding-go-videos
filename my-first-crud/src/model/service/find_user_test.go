package service

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"github.com/arturbaccarin/go-my-first-crud/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

// mockgen -source=src/model/repository/user_repository.go -destination=src/test/mocks/user_repository_mock.go -package=mock
// mockgen -source=src/model/user_domain_interface.go -destination=src/test/mocks/user_domain_interface_mock.go -package=mock

func TestUserDomainService_FindUserByIdServices(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "name", 98)
		userDomain.SetId(id)

		repository.EXPECT().FindUserById(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserById(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetId(), id)
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().FindUserById(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserById(id)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailServices(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@test.com"

		userDomain := model.NewUserDomain(email, "test", "name", 98)
		userDomain.SetId(id)

		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByEmail(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetId(), id)
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@test.com"

		repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByEmail(email)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailAndPasswordServices(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockUserRepository(ctrl)
	service := &userDomainService{
		userRepository: repository}

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := model.NewUserDomain(email, password, "name", 98)
		userDomain.SetId(id)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		userDomainReturn, err := service.findUserByEmailAndPassword(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetId(), id)
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.findUserByEmailAndPassword(email, password)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
	})
}
