package response

type UserResponse struct {
    ID       string          `json:"id"`
    Email    string          `json:"email"`
    Name     string          `json:"name"`
    UserType string          `json:"user_type"`
    WorkInfo *WorkInfoResponse `json:"work_info,omitempty"`
}

type WorkInfoResponse struct {
    Team          string `json:"team"`
    Position      string `json:"position"`
    DefaultShift  string `json:"default_shift"`
    WeekdayOff    string `json:"weekday_off"`
    WeekendDayOff string `json:"weekend_day_off"`
    SuperiorID    string `json:"superior_id"`
}