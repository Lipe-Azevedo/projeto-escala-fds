package model

type Team string

const (
	CustomerService  Team = "Customer Service"
	Security         Team = "security"
	TechnicalSupport Team = "Technical Support"
)

type Shift string

const (
	MorningShift   Shift = "06:00-14:00"
	AfternoonShift Shift = "14:00-22:00"
	NightShift     Shift = "22:00-06:00"
)

type Weekday string

const (
	Monday    Weekday = "monday"
	Tuesday   Weekday = "tuesday"
	Wednesday Weekday = "wednesday"
	Thursday  Weekday = "thursday"
	Friday    Weekday = "friday"
)

type WeekendDayOff string

const (
	Saturday WeekendDayOff = "saturday"
	Sunday   WeekendDayOff = "sunday"
)

type WorkInfoDomain struct {
	userID        string
	team          Team
	position      string
	defaultShift  Shift
	weekdayOff    Weekday
	weekendDayOff WeekendDayOff
	superiorID    string
}

func (w *WorkInfoDomain) GetUserId() string               { return w.userID }
func (w *WorkInfoDomain) GetTeam() Team                   { return w.team }
func (w *WorkInfoDomain) GetPosition() string             { return w.position }
func (w *WorkInfoDomain) GetDefaultShift() Shift          { return w.defaultShift }
func (w *WorkInfoDomain) GetWeekdayOff() Weekday          { return w.weekdayOff }
func (w *WorkInfoDomain) GetWeekendDayOff() WeekendDayOff { return w.weekendDayOff }
func (w *WorkInfoDomain) GetSuperiorID() string           { return w.superiorID }
