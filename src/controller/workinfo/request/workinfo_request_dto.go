package request // O nome do pacote permanece 'request' dentro deste subdiretório

type WorkInfoRequest struct {
	Team          string `json:"team" binding:"required,oneof='Customer Service' security 'Technical Support'"`
	Position      string `json:"position" binding:"required"`
	DefaultShift  string `json:"default_shift" binding:"required,oneof=06:00-14:00 14:00-22:00 22:00-06:00"`
	WeekdayOff    string `json:"weekday_off" binding:"required,oneof=monday tuesday wednesday thursday friday"`
	WeekendDayOff string `json:"weekend_day_off" binding:"required,oneof=saturday sunday"`
	SuperiorID    string `json:"superior_id" binding:"required"` // Mantido como string, validação de existência no serviço
}

type WorkInfoUpdateRequest struct {
	Team          *string `json:"team,omitempty" binding:"omitempty,oneof='Customer Service' security 'Technical Support'"`
	Position      *string `json:"position,omitempty" binding:"omitempty"`
	DefaultShift  *string `json:"default_shift,omitempty" binding:"omitempty,oneof=06:00-14:00 14:00-22:00 22:00-06:00"`
	WeekdayOff    *string `json:"weekday_off,omitempty" binding:"omitempty,oneof=monday tuesday wednesday thursday friday"`
	WeekendDayOff *string `json:"weekend_day_off,omitempty" binding:"omitempty,oneof=saturday sunday"`
	SuperiorID    *string `json:"superior_id,omitempty" binding:"omitempty"` // Mantido como *string
}