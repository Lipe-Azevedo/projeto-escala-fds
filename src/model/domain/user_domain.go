package domain

import (
	"crypto/md5"
	"encoding/hex"
	// "log" // Não precisamos de log aqui se a versão MD5 não falha
)

// UserDomainInterface define a interface para o domínio do usuário.
type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetUserType() UserType

	SetID(string)
	EncryptPassword() error // <<< Assinatura atualizada para retornar erro
	// CheckPassword(plainPassword string) bool // Será adicionado com bcrypt
}

// userDomain é a struct que representa o domínio do usuário.
type userDomain struct {
	id       string
	email    string
	password string
	name     string
	userType UserType
}

// NewUserDomain construtor para criar uma nova instância de UserDomainInterface.
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

// NewUserUpdateDomain construtor para criar um UserDomain para atualização (nome/senha).
func NewUserUpdateDomain(
	name string,
	password string,
) UserDomainInterface {
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

func (ud *userDomain) SetID(id string) { ud.id = id }

// EncryptPassword atualmente usa MD5. Agora retorna nil como erro.
func (ud *userDomain) EncryptPassword() error { // <<< Assinatura atualizada
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
	return nil // <<< Retorna nil, pois MD5 aqui não falha
}
