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

func TestUserRepository_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collectionName)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_user_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowldged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)

		domain := model.NewUserDomain("test@test.com", "test", "name", 98)
		domain.SetId(primitive.NewObjectID().Hex())

		err := repo.UpdateUser(domain.GetId(), domain)

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)

		domain := model.NewUserDomain("test@test.com", "test", "name", 98)
		domain.SetId(primitive.NewObjectID().Hex())

		err := repo.UpdateUser(domain.GetId(), domain)

		assert.NotNil(t, err)
	})
}
