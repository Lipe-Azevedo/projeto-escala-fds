package model

import "time"

type SwapDomainInterface interface {
	GetID() string
	GetRequesterID() string
	GetRequestedID() string
	GetCurrentShift() Shift
	GetNewShift() Shift
	GetCurrentDayOff() Weekday
	GetNewDayOff() Weekday
	GetStatus() SwapStatus
	GetReason() string
	GetCreatedAt() time.Time
	GetApprovedAt() *time.Time
	GetApprovedBy() *string

	SetID(string)
	SetStatus(SwapStatus)
	SetApprovedAt(time.Time)
	SetApprovedBy(string)
}

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
		status:        StatusPending,
		reason:        reason,
		createdAt:     time.Now(),
	}
}

func NewSwapUpdateDomain(
	requestedID string,
	currentShift Shift,
	newShift Shift,
	currentDayOff Weekday,
	newDayOff Weekday,
	reason string,
) SwapDomainInterface {
	return &swapDomain{
		requestedID:   requestedID,
		currentShift:  currentShift,
		newShift:      newShift,
		currentDayOff: currentDayOff,
		newDayOff:     newDayOff,
		reason:        reason,
		createdAt:     time.Now(),
	}
}
