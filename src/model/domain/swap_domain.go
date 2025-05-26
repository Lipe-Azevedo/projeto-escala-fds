package domain

import "time"

// SwapDomainInterface define a interface para o domínio de Swap.
type SwapDomainInterface interface {
	GetID() string
	GetRequesterID() string
	GetRequestedID() string
	GetCurrentShift() Shift    // Vem de common_types.go
	GetNewShift() Shift        // Vem de common_types.go
	GetCurrentDayOff() Weekday // Vem de common_types.go
	GetNewDayOff() Weekday     // Vem de common_types.go
	GetStatus() SwapStatus     // Vem de common_types.go
	GetReason() string
	GetCreatedAt() time.Time
	GetApprovedAt() *time.Time
	GetApprovedBy() *string

	SetID(string)
	SetStatus(SwapStatus)
	SetApprovedAt(approvedAt time.Time) // Mantido como time.Time, a implementação usa ponteiro
	SetApprovedBy(approvedBy string)    // Mantido como string, a implementação usa ponteiro

	// Novos setters para alinhar com a lógica de update_swap_service
	SetRequesterID(id string)   // Adicionado se necessário
	SetRequestedID(id string)   // Adicionado
	SetCurrentShift(s Shift)    // Adicionado
	SetNewShift(s Shift)        // Adicionado
	SetCurrentDayOff(d Weekday) // Adicionado
	SetNewDayOff(d Weekday)     // Adicionado
	SetReason(r string)         // Adicionado
	// ClearApprovedAt() // Necessário se quisermos explicitamente anular via interface
	// ClearApprovedBy() // Necessário se quisermos explicitamente anular via interface
}

// swapDomain é a struct que representa o domínio de Swap.
type swapDomain struct {
	id            string
	requesterID   string
	requestedID   string
	currentShift  Shift
	newShift      Shift
	currentDayOff Weekday
	newDayOff     Weekday
	status        SwapStatus
	reason        string
	createdAt     time.Time
	approvedAt    *time.Time
	approvedBy    *string
}

// NewSwapDomain construtor para SwapDomainInterface.
func NewSwapDomain(
	requesterID string,
	requestedID string,
	currentShift Shift,
	newShift Shift,
	currentDayOff Weekday,
	newDayOff Weekday,
	reason string,
) SwapDomainInterface {
	return &swapDomain{
		requesterID:   requesterID,
		requestedID:   requestedID,
		currentShift:  currentShift,
		newShift:      newShift,
		currentDayOff: currentDayOff,
		newDayOff:     newDayOff,
		status:        StatusPending, // StatusPending de common_types.go
		reason:        reason,
		createdAt:     time.Now(),
	}
}

// NewSwapUpdateDomain foi usado no seu código original.
// Ele é similar ao NewSwapDomain mas não inclui requesterID e status inicial.
// Pode ser usado para construir um domain com os campos atualizáveis.
func NewSwapUpdateDomain(
	requestedID string,
	currentShift Shift,
	newShift Shift,
	currentDayOff Weekday,
	newDayOff Weekday,
	reason string,
) SwapDomainInterface {
	return &swapDomain{
		// requesterID é omitido
		requestedID:   requestedID,
		currentShift:  currentShift,
		newShift:      newShift,
		currentDayOff: currentDayOff,
		newDayOff:     newDayOff,
		// status é omitido (será setado depois)
		reason: reason,
		// createdAt é omitido (não deve ser alterado em um update comum)
		// approvedAt e approvedBy são omitidos (serão setados depois)
	}
}

func (s *swapDomain) GetID() string             { return s.id }
func (s *swapDomain) GetRequesterID() string    { return s.requesterID }
func (s *swapDomain) GetRequestedID() string    { return s.requestedID }
func (s *swapDomain) GetCurrentShift() Shift    { return s.currentShift }
func (s *swapDomain) GetNewShift() Shift        { return s.newShift }
func (s *swapDomain) GetCurrentDayOff() Weekday { return s.currentDayOff }
func (s *swapDomain) GetNewDayOff() Weekday     { return s.newDayOff }
func (s *swapDomain) GetStatus() SwapStatus     { return s.status }
func (s *swapDomain) GetReason() string         { return s.reason }
func (s *swapDomain) GetCreatedAt() time.Time   { return s.createdAt }
func (s *swapDomain) GetApprovedAt() *time.Time { return s.approvedAt }
func (s *swapDomain) GetApprovedBy() *string    { return s.approvedBy }

func (s *swapDomain) SetID(id string)                    { s.id = id }
func (s *swapDomain) SetStatus(status SwapStatus)        { s.status = status }
func (s *swapDomain) SetApprovedAt(approvedAt time.Time) { s.approvedAt = &approvedAt }
func (s *swapDomain) SetApprovedBy(approvedBy string)    { s.approvedBy = &approvedBy }

// Implementações para os novos setters (se decidirmos adicioná-los à interface)
func (s *swapDomain) SetRequesterID(id string)   { s.requesterID = id }
func (s *swapDomain) SetRequestedID(id string)   { s.requestedID = id }
func (s *swapDomain) SetCurrentShift(sh Shift)   { s.currentShift = sh }
func (s *swapDomain) SetNewShift(sh Shift)       { s.newShift = sh }
func (s *swapDomain) SetCurrentDayOff(d Weekday) { s.currentDayOff = d }
func (s *swapDomain) SetNewDayOff(d Weekday)     { s.newDayOff = d }
func (s *swapDomain) SetReason(r string)         { s.reason = r }

// Para limpar ApprovedAt/By, precisaríamos de métodos como:
// func (s *swapDomain) ClearApprovedAt() { s.approvedAt = nil }
// func (s *swapDomain) ClearApprovedBy() { s.approvedBy = nil }
// E adicioná-los à interface.
