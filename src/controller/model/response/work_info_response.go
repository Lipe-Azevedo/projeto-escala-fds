package response

type WorkInfoResponse struct {
	Team          string `json:"team"`
	Position      string `json:"position"`
	DefaultShift  string `json:"default_shift"`
	WeekdayOff    string `json:"weekday_off"`
	WeekendDayOff string `json:"weekend_day_off"`
	SuperiorID    string `json:"superior_id"`
}
