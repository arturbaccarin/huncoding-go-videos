package controller

import (
	"fmt"
	"net/http"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/validation"
	"github.com/arturbaccarin/go-my-first-crud/src/controller/model/request"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"github.com/arturbaccarin/go-my-first-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller", zap.String("journey", "LoginUser"))

	var userRequest request.UserLogin
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "LoginUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)

	domainResult, token, err := uc.service.LoginUser(domain)
	if err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "LoginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("LoginUser controller executed successfully!", zap.String("journey", "LoginUser"))

	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
