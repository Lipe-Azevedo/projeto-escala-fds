package domain

import (
	"crypto/md5" // Temporário, será substituído por bcrypt
	"encoding/hex"
)

// UserDomainInterface define a interface para o domínio do usuário.
type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetUserType() UserType // UserType vem de common_types.go (mesmo pacote 'domain')

	SetID(string)
	EncryptPassword() // Temporário, retornará erro com bcrypt
	// CheckPassword(plainPassword string) bool // Será adicionado com bcrypt
}

// userDomain é a struct que representa o domínio do usuário.
type userDomain struct {
	id       string
	email    string
	password string
	name     string
	userType UserType // UserType vem de common_types.go (mesmo pacote 'domain')
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

// EncryptPassword atualmente usa MD5. Será refatorado para bcrypt.
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
