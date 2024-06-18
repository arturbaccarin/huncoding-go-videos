package controller

import (
	"net/http"
	"net/mail"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller", zap.String("journey", "FindUserById"))

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		logger.Error("Error FindUserById controller", err, zap.String("journey", "FindUserById"))

		errorMessage := rest_err.NewBadRequestError("id is invalid")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserById(id)
	if err != nil {
		logger.Error("Error FindUserById controller", err, zap.String("journey", "FindUserById"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("Error FindUserById controller", zap.String("journey", "FindUserById"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller", zap.String("journey", "FindUserByEmail"))

	email := c.Param("email")

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Error FindUserByEmail controller", err, zap.String("journey", "FindUserByEmail"))

		errorMessage := rest_err.NewBadRequestError("email is invalid")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error FindUserByEmail controller", err, zap.String("journey", "FindUserByEmail"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("Error FindUserByEmail controller", zap.String("journey", "FindUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
