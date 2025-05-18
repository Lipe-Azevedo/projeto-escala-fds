package response

type UserResponse struct {
	ID       string            `json:"id"`
	Email    string            `json:"email"`
	Name     string            `json:"name"`
	UserType string            `json:"user_type"`
	WorkInfo *WorkInfoResponse `json:"work_info,omitempty"`
}
