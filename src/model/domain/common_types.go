package domain

// UserType define o tipo de usuário no sistema.
type UserType string

const (
	UserTypeCollaborator UserType = "colaborador"
	UserTypeMaster       UserType = "master"
)

// Team define as equipes de trabalho.
type Team string

const (
	TeamCustomerService  Team = "Customer Service"
	TeamSecurity         Team = "Security"
	TeamTechnicalSupport Team = "Technical Support"
)

// Shift define os turnos de trabalho.
type Shift string

const (
	ShiftMorning   Shift = "06:00-14:00"
	ShiftAfternoon Shift = "14:00-22:00"
	ShiftNight     Shift = "22:00-06:00"
)

// Weekday define os dias da semana para folgas.
type Weekday string

const (
	WeekdayMonday    Weekday = "monday"
	WeekdayTuesday   Weekday = "tuesday"
	WeekdayWednesday Weekday = "wednesday"
	WeekdayThursday  Weekday = "thursday"
	WeekdayFriday    Weekday = "friday"
)

// WeekendDayOff define os dias do fim de semana para folgas.
type WeekendDayOff string

const (
	WeekendSaturday WeekendDayOff = "saturday"
	WeekendSunday   WeekendDayOff = "sunday"
)

// SwapStatus define os status de uma solicitação de troca.
type SwapStatus string

const (
	StatusPending  SwapStatus = "pending"
	StatusApproved SwapStatus = "approved"
	StatusRejected SwapStatus = "rejected"
)
