package request // O nome do pacote permanece 'request' dentro deste subdiretório

// SwapRequest é usado para criar uma nova solicitação de troca
// e também para atualizar o status (onde alguns campos podem não ser relevantes para o update de status,
// mas o binding "oneof" para Status ainda é útil).
type SwapRequest struct {
	// Campos para criação de Swap
	RequestedID   string `json:"requested_id,omitempty" binding:"required_without=Status"` // Obrigatório para criar, opcional para update de status
	CurrentShift  string `json:"current_shift,omitempty" binding:"required_without=Status,omitempty,oneof=06:00-14:00 14:00-22:00 22:00-06:00"`
	NewShift      string `json:"new_shift,omitempty" binding:"required_without=Status,omitempty,oneof=06:00-14:00 14:00-22:00 22:00-06:00"`
	CurrentDayOff string `json:"current_day_off,omitempty" binding:"required_without=Status,omitempty,oneof=monday tuesday wednesday thursday friday"`
	NewDayOff     string `json:"new_day_off,omitempty" binding:"required_without=Status,omitempty,oneof=monday tuesday wednesday thursday friday"`
	Reason        string `json:"reason,omitempty" binding:"max=500"`

	// Campo para atualização de status
	Status string `json:"status,omitempty" binding:"required_without=RequestedID,omitempty,oneof=pending approved rejected"`
}

// NOTA: O SwapRequest original tinha todos os campos como 'required'.
// Para o endpoint de atualização de status (PUT /:id/status), geralmente apenas o campo 'status' é enviado.
// Para o endpoint de criação (POST /shift-swap), os outros campos são necessários.
// Ajustei os bindings para tentar acomodar ambos, usando `required_without`.
// Uma alternativa seria ter DTOs separados: SwapCreationRequest e SwapStatusUpdateRequest.
// Por ora, mantive um DTO único com bindings condicionais.
