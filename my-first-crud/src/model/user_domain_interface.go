package model

import "github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	SetPassword(string)
	GetPassword() string
	GetAge() int8
	GetName() string

	GetId() string
	SetId(string)

	EncryptPassword()

	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(name string, age int8) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
