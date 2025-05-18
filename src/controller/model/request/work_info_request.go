package request

type WorkInfoRequest struct {
	Team          string `json:"team" binding:"required"`
	Position      string `json:"position" binding:"required"`
	DefaultShift  string `json:"default_shift" binding:"required,oneof=06:00-14:00 14:00-22:00 22:00-06:00"`
	WeekdayOff    string `json:"weekday_off" binding:"required,oneof=monday tuesday wednesday thursday friday"`
	WeekendDayOff string `json:"weekend_day_off" binding:"required,oneof=saturday sunday"`
	SuperiorID    string `json:"superior_id" binding:"required"`
}
