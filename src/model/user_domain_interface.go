package model

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int

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
		age:      age,
	}
}

func NewUserUpdateDomain(
	name string,
	age int,

) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}
