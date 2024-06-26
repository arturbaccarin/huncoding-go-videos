package service

import (
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init createUser model", zap.String("journey", "createUser"))

	user, _ := ud.userRepository.FindUserByEmail(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("email is already registered in another account")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("init createUser model", err, zap.String("journey", "createUser"))
		return nil, err
	}

	return userDomainRepository, nil
}
