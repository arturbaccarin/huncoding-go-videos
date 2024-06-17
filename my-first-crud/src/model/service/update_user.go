package service

import (
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("init UpdateUser model", zap.String("journey", "UpdateUser"))

	return nil
}
