package model

type UserType string

const (
	UserTypeCollaborator UserType = "colaborador"
	UserTypeMaster       UserType = "master"
)

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	userType UserType
}

func (ud *userDomain) GetID() string         { return ud.id }
func (ud *userDomain) GetEmail() string      { return ud.email }
func (ud *userDomain) GetPassword() string   { return ud.password }
func (ud *userDomain) GetName() string       { return ud.name }
func (ud *userDomain) GetUserType() UserType { return ud.userType }

func (ud *userDomain) SetID(id string)               { ud.id = id }
func (ud *userDomain) SetEmail(email string)         { ud.email = email }
func (ud *userDomain) SetPassword(password string)   { ud.password = password }
func (ud *userDomain) SetName(name string)           { ud.name = name }
func (ud *userDomain) SetUserType(userType UserType) { ud.userType = userType }
