package model

import "time"

type ShiftSwapStatus string

const (
	StatusPending  ShiftSwapStatus = "pending"
	StatusApproved ShiftSwapStatus = "approved"
	StatusRejected ShiftSwapStatus = "rejected"
)

type shiftSwapDomain struct {
	id            string
	requesterID   string
	requestedID   string
	currentShift  Shift
	newShift      Shift
	currentDayOff Weekday
	newDayOff     Weekday
	status        ShiftSwapStatus
	reason        string
	createdAt     time.Time
	approvedAt    *time.Time
	approvedBy    *string
}

func (s *shiftSwapDomain) GetID() string              { return s.id }
func (s *shiftSwapDomain) GetRequesterID() string     { return s.requesterID }
func (s *shiftSwapDomain) GetRequestedID() string     { return s.requestedID }
func (s *shiftSwapDomain) GetCurrentShift() Shift     { return s.currentShift }
func (s *shiftSwapDomain) GetNewShift() Shift         { return s.newShift }
func (s *shiftSwapDomain) GetCurrentDayOff() Weekday  { return s.currentDayOff }
func (s *shiftSwapDomain) GetNewDayOff() Weekday      { return s.newDayOff }
func (s *shiftSwapDomain) GetStatus() ShiftSwapStatus { return s.status }
func (s *shiftSwapDomain) GetReason() string          { return s.reason }
func (s *shiftSwapDomain) GetCreatedAt() time.Time    { return s.createdAt }
func (s *shiftSwapDomain) GetApprovedAt() *time.Time  { return s.approvedAt }
func (s *shiftSwapDomain) GetApprovedBy() *string     { return s.approvedBy }

func (s *shiftSwapDomain) SetID(id string)                    { s.id = id }
func (s *shiftSwapDomain) SetStatus(status ShiftSwapStatus)   { s.status = status }
func (s *shiftSwapDomain) SetApprovedAt(approvedAt time.Time) { s.approvedAt = &approvedAt }
func (s *shiftSwapDomain) SetApprovedBy(approvedBy string)    { s.approvedBy = &approvedBy }
