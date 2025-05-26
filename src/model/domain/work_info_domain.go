package domain

// WorkInfoDomainInterface define a interface para o domínio de WorkInfo.
type WorkInfoDomainInterface interface {
	GetUserId() string
	GetTeam() Team // Vem de common_types.go
	GetPosition() string
	GetDefaultShift() Shift          // Vem de common_types.go
	GetWeekdayOff() Weekday          // Vem de common_types.go
	GetWeekendDayOff() WeekendDayOff // Vem de common_types.go
	GetSuperiorID() string

	SetTeam(team Team)
	SetPosition(position string)
	SetDefaultShift(shift Shift)
	SetWeekdayOff(day Weekday)
	SetWeekendDayOff(day WeekendDayOff)
	SetSuperiorID(id string)
}

// WorkInfoDomain é a struct que representa o domínio de WorkInfo.
type workInfoDomain struct { // Renomeado de WorkInfoDomain para workInfoDomain (minúsculo) para consistência
	userID        string
	team          Team
	position      string
	defaultShift  Shift
	weekdayOff    Weekday
	weekendDayOff WeekendDayOff
	superiorID    string
}

// NewWorkInfoDomain construtor para WorkInfoDomainInterface.
func NewWorkInfoDomain(
	userID string,
	team Team,
	position string,
	defaultShift Shift,
	weekdayOff Weekday,
	weekendDayOff WeekendDayOff,
	superiorID string,
) WorkInfoDomainInterface {
	return &workInfoDomain{
		userID:        userID,
		team:          team,
		position:      position,
		defaultShift:  defaultShift,
		weekdayOff:    weekdayOff,
		weekendDayOff: weekendDayOff,
		superiorID:    superiorID,
	}
}

// NewWorkInfoUpdateDomain (se necessário para uma atualização mais granular, não presente no seu código original para domain)
// O seu NewWorkInfoUpdateDomain original criava uma instância de WorkInfoDomain, o que é um pouco confuso.
// Geralmente, para atualizações, você carregaria o domínio existente e usaria setters,
// ou um DTO de atualização específico seria mapeado para os setters.
// Vou omitir NewWorkInfoUpdateDomain por enquanto, pois os setters já existem.
// O seu NewWorkInfoUpdateDomain original:
/*
func NewWorkInfoUpdateDomain(
    team Team,
    position string,
    defaultShift Shift,
    weekdayOff Weekday,
    superiorID string,
) WorkInfoDomainInterface {
    return &workInfoDomain{ // Note que faltava weekendDayOff aqui e userID
        team:         team,
        position:     position,
        defaultShift: defaultShift,
        weekdayOff:   weekdayOff,
        superiorID:   superiorID,
    }
}
*/

func (w *workInfoDomain) GetUserId() string               { return w.userID }
func (w *workInfoDomain) GetTeam() Team                   { return w.team }
func (w *workInfoDomain) GetPosition() string             { return w.position }
func (w *workInfoDomain) GetDefaultShift() Shift          { return w.defaultShift }
func (w *workInfoDomain) GetWeekdayOff() Weekday          { return w.weekdayOff }
func (w *workInfoDomain) GetWeekendDayOff() WeekendDayOff { return w.weekendDayOff }
func (w *workInfoDomain) GetSuperiorID() string           { return w.superiorID }

func (w *workInfoDomain) SetTeam(team Team)                  { w.team = team }
func (w *workInfoDomain) SetPosition(position string)        { w.position = position }
func (w *workInfoDomain) SetDefaultShift(shift Shift)        { w.defaultShift = shift }
func (w *workInfoDomain) SetWeekdayOff(day Weekday)          { w.weekdayOff = day }
func (w *workInfoDomain) SetWeekendDayOff(day WeekendDayOff) { w.weekendDayOff = day }
func (w *workInfoDomain) SetSuperiorID(id string)            { w.superiorID = id }
