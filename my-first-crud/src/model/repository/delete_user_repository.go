package repository

import (
	"context"
	"os"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(id string) *rest_err.RestErr {

	logger.Info("Init DeleteUser repository", zap.String("journey", "DeleteUser"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("Init DeleteUser succes!", zap.String("userId", objectId.Hex()), zap.String("journey", "DeleteUser"))

	return nil
}
