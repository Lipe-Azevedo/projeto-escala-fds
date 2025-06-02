package domain

type UserType string

const (
	UserTypeCollaborator UserType = "colaborador"
	UserTypeMaster       UserType = "master"
)

type Team string

const (
	TeamCustomerService  Team = "Customer Service"
	TeamSecurity         Team = "Security"
	TeamTechnicalSupport Team = "Technical Support"
)

type Shift string

const (
	ShiftMorning   Shift = "06:00-14:00"
	ShiftAfternoon Shift = "14:00-22:00"
	ShiftNight     Shift = "22:00-06:00"
)

type Weekday string

const (
	WeekdayMonday    Weekday = "monday"
	WeekdayTuesday   Weekday = "tuesday"
	WeekdayWednesday Weekday = "wednesday"
	WeekdayThursday  Weekday = "thursday"
	WeekdayFriday    Weekday = "friday"
)

type WeekendDayOff string

const (
	WeekendSaturday WeekendDayOff = "saturday"
	WeekendSunday   WeekendDayOff = "sunday"
)

type SwapStatus string

const (
	StatusPending  SwapStatus = "pending"
	StatusApproved SwapStatus = "approved"
	StatusRejected SwapStatus = "rejected"
)
