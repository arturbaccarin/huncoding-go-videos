package repository

import (
	"os"
	"testing"

	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {
	databaseName := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_domain_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowldged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.CreateUser(model.NewUserDomain("test@test.com", "test", "name", 98))
		_, errId := primitive.ObjectIDFromHex(userDomain.GetId())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
		assert.EqualValues(t, userDomain.GetName(), "name")
		assert.EqualValues(t, userDomain.GetAge(), 98)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		_, err := repo.CreateUser(model.NewUserDomain("test@test.com", "test", "name", 98))

		assert.NotNil(t, err)
	})
}
