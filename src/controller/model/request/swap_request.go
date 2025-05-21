package request

type ShiftSwapRequest struct {
	RequestedID   string `json:"requested_id" binding:"required"`
	CurrentShift  string `json:"current_shift" binding:"required,oneof=06:00-14:00 14:00-22:00 22:00-06:00"`
	NewShift      string `json:"new_shift" binding:"required,oneof=06:00-14:00 14:00-22:00 22:00-06:00"`
	CurrentDayOff string `json:"current_day_off" binding:"required,oneof=monday tuesday wednesday thursday friday"`
	NewDayOff     string `json:"new_day_off" binding:"required,oneof=monday tuesday wednesday thursday friday"`
	Status        string `json:"status" binding:"required,oneof=pending approved rejected"`
	Reason        string `json:"reason" binding:"max=500"`
}
