package domain

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetUserType() UserType

	SetID(string)
	EncryptPassword() error
	CheckPassword(plainPassword string) bool
}

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	userType UserType
	// workInfo WorkInfoDomainInterface // REMOVIDO
}

func NewUserDomain(email string, password string, name string, userType UserType) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		userType: userType,
	}
}

func NewUserUpdateDomain(name string, password string) UserDomainInterface {
	return &userDomain{
		name:     name,
		password: password,
	}
}

func (ud *userDomain) GetID() string         { return ud.id }
func (ud *userDomain) GetEmail() string      { return ud.email }
func (ud *userDomain) GetPassword() string   { return ud.password }
func (ud *userDomain) GetName() string       { return ud.name }
func (ud *userDomain) GetUserType() UserType { return ud.userType }

// func (ud *userDomain) GetWorkInfo() WorkInfoDomainInterface { return ud.workInfo } // REMOVIDO

func (ud *userDomain) SetID(id string) { ud.id = id }

// func (ud *userDomain) SetWorkInfo(wi WorkInfoDomainInterface) { ud.workInfo = wi } // REMOVIDO

// ... EncryptPassword e CheckPassword permanecem ...
func (ud *userDomain) EncryptPassword() error {
	if ud.password == "" {
		return nil
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating bcrypt hash for user %s: %v", ud.email, err)
		return err
	}
	ud.password = string(hashedPassword)
	return nil
}

func (ud *userDomain) CheckPassword(plainPassword string) bool {
	if ud.password == "" || plainPassword == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(ud.password), []byte(plainPassword))
	return err == nil
}
