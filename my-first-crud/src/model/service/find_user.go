package service

import (
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail service", zap.String("journey", "FindUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserById service", zap.String("journey", "FindUserById"))
	return ud.userRepository.FindUserById(id)
}
