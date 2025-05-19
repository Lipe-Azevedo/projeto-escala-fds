package model

type WorkInfoDomainInterface interface {
	GetUserId() string
	GetTeam() Team
	GetPosition() string
	GetDefaultShift() Shift
	GetWeekdayOff() Weekday
	GetWeekendDayOff() WeekendDayOff
	GetSuperiorID() string

	SetTeam(team Team)
	SetPosition(position string)
	SetDefaultShift(shift Shift)
	SetWeekdayOff(day Weekday)
	SetWeekendDayOff(day WeekendDayOff)
	SetSuperiorID(id string)
}

func NewWorkInfoDomain(
	userID string,
	team Team,
	position string,
	defaultShift Shift,
	weekdayOff Weekday,
	weekendDayOff WeekendDayOff,
	superiorID string,

) WorkInfoDomainInterface {
	return &WorkInfoDomain{
		userID:        userID,
		team:          team,
		position:      position,
		defaultShift:  defaultShift,
		weekdayOff:    weekdayOff,
		weekendDayOff: weekendDayOff,
		superiorID:    superiorID,
	}
}

func NewWorkUpdateDomain(
	team Team,
	position string,
	defaultShift Shift,
	weekdayOff Weekday,
	superiorID string,

) WorkInfoDomainInterface {
	return &WorkInfoDomain{
		team:         team,
		position:     position,
		defaultShift: defaultShift,
		weekdayOff:   weekdayOff,
		superiorID:   superiorID,
	}
}
