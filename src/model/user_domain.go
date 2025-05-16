package model

type UserType string

const (
	UserTypeCollaborator UserType = "colaborador"
	UserTypeMaster       UserType = "master"
)

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetUserType() UserType
	GetWorkInfo() *WorkInfo

	SetID(string)
	SetWorkInfo(*WorkInfo)

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	userType UserType,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		userType: userType,
	}
}

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	userType UserType
	workInfo *WorkInfo
}

func (ud *userDomain) GetID() string              { return ud.id }
func (ud *userDomain) SetID(id string)            { ud.id = id }
func (ud *userDomain) GetEmail() string           { return ud.email }
func (ud *userDomain) GetPassword() string        { return ud.password }
func (ud *userDomain) GetName() string            { return ud.name }
func (ud *userDomain) GetUserType() UserType      { return ud.userType }
func (ud *userDomain) GetWorkInfo() *WorkInfo     { return ud.workInfo }
func (ud *userDomain) SetWorkInfo(info *WorkInfo) { ud.workInfo = info }
