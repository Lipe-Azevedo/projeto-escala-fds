package model

import "time"

type SwapStatus string

const (
	StatusPending  SwapStatus = "pending"
	StatusApproved SwapStatus = "approved"
	StatusRejected SwapStatus = "rejected"
)

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
