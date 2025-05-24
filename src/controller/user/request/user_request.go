package request // Pacote alterado

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	UserType string `json:"user_type" binding:"required,oneof=colaborador master"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"omitempty,min=4,max=100"`
	Password string `json:"password" binding:"omitempty,min=6,containsany=!@#$%*"`
}
