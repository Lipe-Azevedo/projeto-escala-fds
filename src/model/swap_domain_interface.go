package model

import "time"

type ShiftSwapDomainInterface interface {
	GetID() string
	GetRequesterID() string
	GetRequestedID() string
	GetCurrentShift() Shift
	GetNewShift() Shift
	GetCurrentDayOff() Weekday
	GetNewDayOff() Weekday
	GetStatus() ShiftSwapStatus
	GetReason() string
	GetCreatedAt() time.Time
	GetApprovedAt() *time.Time
	GetApprovedBy() *string

	SetID(string)
	SetStatus(ShiftSwapStatus)
	SetApprovedAt(time.Time)
	SetApprovedBy(string)
}

func NewShiftSwapDomain(
	requesterID string,
	requestedID string,
	currentShift Shift,
	newShift Shift,
	currentDayOff Weekday,
	newDayOff Weekday,
	reason string,
) ShiftSwapDomainInterface {
	return &shiftSwapDomain{
		requesterID:   requesterID,
		requestedID:   requestedID,
		currentShift:  currentShift,
		newShift:      newShift,
		currentDayOff: currentDayOff,
		newDayOff:     newDayOff,
		status:        StatusPending,
		reason:        reason,
		createdAt:     time.Now(),
	}
}

func NewShiftSwapUpdateDomain(
	requestedID string,
	currentShift Shift,
	newShift Shift,
	currentDayOff Weekday,
	newDayOff Weekday,
	reason string,
) ShiftSwapDomainInterface {
	return &shiftSwapDomain{
		requestedID:   requestedID,
		currentShift:  currentShift,
		newShift:      newShift,
		currentDayOff: currentDayOff,
		newDayOff:     newDayOff,
		reason:        reason,
		createdAt:     time.Now(),
	}
}
