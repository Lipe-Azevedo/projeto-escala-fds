package model

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetUserType() UserType
	SetID(string)

	EncryptPassword()
}

func NewUserDomain(
	email string,
	password string,
	name string,
	userType UserType,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		userType: userType,
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
