package service

import (
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init UpdateUser model", zap.String("journey", "UpdateUser"))

	err := ud.userRepository.UpdateUser(id, userDomain)
	if err != nil {
		logger.Error("init UpdateUser model", err, zap.String("journey", "UpdateUser"))
		return err
	}

	return nil
}
