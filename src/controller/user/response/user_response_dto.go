package response

import (
	// Import para o DTO de WorkInfoResponse (já deve estar usando o caminho correto após reorganização do WorkInfo)
	workinfo_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/response"
)

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
	// Descomentado e tipo ajustado para usar o DTO específico de WorkInfo
	WorkInfo *workinfo_response_dto.WorkInfoResponse `json:"work_info,omitempty"`
}
