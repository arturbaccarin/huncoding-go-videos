package controller

import (
	"net/http"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/validation"
	"github.com/arturbaccarin/go-my-first-crud/src/controller/model/request"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUserById(c *gin.Context) {
	logger.Info("Init UpdateUserById controller", zap.String("journey", "UpdateUserById"))

	var userRequest request.UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "UpdateUserById"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	id := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		logger.Error("Error UpdateUserById controller", err, zap.String("journey", "UpdateUserById"))

		errorMessage := rest_err.NewBadRequestError("id is invalid")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateUser(id, domain)
	if err != nil {
		logger.Error("Error trying to call updateUser", err, zap.String("journey", "UpdateUserById"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated!", zap.String("journey", "UpdateUserById"))

	c.Status(http.StatusNoContent)
}
