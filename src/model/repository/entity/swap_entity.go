package entity

import "time"

type SwapEntity struct {
	ID            string     `bson:"_id,omitempty"`
	RequesterID   string     `bson:"requester_id"`
	RequestedID   string     `bson:"requested_id"`
	CurrentShift  string     `bson:"current_shift"`
	NewShift      string     `bson:"new_shift"`
	CurrentDayOff string     `bson:"current_day_off"`
	NewDayOff     string     `bson:"new_day_off"`
	Status        string     `bson:"status"`
	Reason        string     `bson:"reason"`
	CreatedAt     time.Time  `bson:"created_at"`
	ApprovedAt    *time.Time `bson:"approved_at,omitempty"`
	ApprovedBy    *string    `bson:"approved_by,omitempty"`
}
