package service

import (
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("init DeleteUser model", zap.String("journey", "DeleteUser"))

	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		logger.Error("init DeleteUser model", err, zap.String("journey", "DeleteUser"))
		return err
	}

	return nil
}
