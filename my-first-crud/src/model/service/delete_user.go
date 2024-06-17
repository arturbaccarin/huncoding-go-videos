package service

import (
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*userDomainService) DeleteUser(string) *rest_err.RestErr {
	logger.Info("init DeleteUser model", zap.String("journey", "deleteUser"))
	return nil
}
