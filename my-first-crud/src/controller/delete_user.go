package controller

import (
	"net/http"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUserById(c *gin.Context) {
	logger.Info("Init DeleteUserById controller", zap.String("journey", "DeleteUserById"))

	id := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		logger.Error("Error DeleteUserById controller", err, zap.String("journey", "DeleteUserById"))

		errorMessage := rest_err.NewBadRequestError("id is invalid")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	err := uc.service.DeleteUser(id)
	if err != nil {
		logger.Error("Error trying to call updateUser", err, zap.String("journey", "DeleteUserById"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted!", zap.String("journey", "DeleteUserById"))

	c.Status(http.StatusNoContent)
}
