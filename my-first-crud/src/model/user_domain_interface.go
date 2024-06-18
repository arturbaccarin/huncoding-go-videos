package model

type UserDomainInterface interface {
	GetEmail() string
	SetPassword(string)
	GetPassword() string
	GetAge() int8
	GetName() string

	GetId() string
	SetId(string)

	EncryptPassword()
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
