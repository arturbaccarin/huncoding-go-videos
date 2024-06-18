package service

import (
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/logger"
	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPassword(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, err
	}

	return user, nil
}
