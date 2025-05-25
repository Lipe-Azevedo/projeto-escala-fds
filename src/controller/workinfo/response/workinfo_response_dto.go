package response // O nome do pacote permanece 'response' dentro deste subdiretório

type WorkInfoResponse struct {
	UserID        string `json:"user_id"` // Este campo corresponde ao _id do WorkInfo no banco, que é o ID do usuário.
	Team          string `json:"team"`
	Position      string `json:"position"`
	DefaultShift  string `json:"default_shift"`
	WeekdayOff    string `json:"weekday_off"`
	WeekendDayOff string `json:"weekend_day_off"`
	SuperiorID    string `json:"superior_id"`
}