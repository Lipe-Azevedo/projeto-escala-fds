package response // Pacote alterado

// ATENÇÃO: WorkInfoResponse será movido para src/controller/workinfo/response/
// Se UserResponse precisar dele, o import precisará ser ajustado quando WorkInfo for reorganizado.
// Por enquanto, vamos manter a referência como está, mas ela quebrará temporariamente
// até que WorkInfoResponse seja movido e o import aqui seja atualizado para
// "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/response"
// Ou, se WorkInfoResponse for um tipo mais genérico, ele pode ficar em um local comum.
// Dado que você quer subpastas, o mais provável é que ele vá para workinfo.
// Vou comentar o campo WorkInfo por enquanto para evitar erro de compilação imediato,
// e o descomentaremos e ajustaremos o import quando chegarmos em WorkInfo.

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
	// WorkInfo *WorkInfoResponse `json:"work_info,omitempty"` // Comentado temporariamente
}
