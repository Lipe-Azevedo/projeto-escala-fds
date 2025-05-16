package model

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string

	SetID(string)

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int,

) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
	}
}

func NewUserUpdateDomain(
	name string,
	password string,

) UserDomainInterface {
	return &userDomain{
		name:     name,
		password: password,
	}
}
