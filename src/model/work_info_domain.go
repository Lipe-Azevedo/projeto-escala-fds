package model

import "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"

// Dias úteis para folga fixa
type Weekday string

const (
	Monday    Weekday = "monday"
	Tuesday   Weekday = "tuesday"
	Wednesday Weekday = "wednesday"
	Thursday  Weekday = "thursday"
	Friday    Weekday = "friday"
)

// Tipos de turno
type Shift string

const (
	MorningShift   Shift = "06:00-14:00"
	AfternoonShift Shift = "14:00-22:00"
	NightShift     Shift = "22:00-06:00"
)

// Folgas de fim de semana
type WeekendDayOff string

const (
	Saturday WeekendDayOff = "saturday"
	Sunday   WeekendDayOff = "sunday"
)

type WorkInfo struct {
	Team          string        `json:"team"`
	Position      string        `json:"position"`
	DefaultShift  Shift         `json:"default_shift"`
	WeekdayOff    Weekday       `json:"weekday_off"`
	WeekendDayOff WeekendDayOff `json:"weekend_day_off"`
	SuperiorID    string        `json:"superior_id"`
}

type WorkInfoInterface interface {
	GetTeam() string
	GetPosition() string
	GetDefaultShift() Shift
	GetWeekdayOff() Weekday
	GetWeekendDayOff() WeekendDayOff
	GetSuperiorID() string

	Validate() *rest_err.RestErr
}

func (w *WorkInfo) GetTeam() string                 { return w.Team }
func (w *WorkInfo) GetPosition() string             { return w.Position }
func (w *WorkInfo) GetDefaultShift() Shift          { return w.DefaultShift }
func (w *WorkInfo) GetWeekdayOff() Weekday          { return w.WeekdayOff }
func (w *WorkInfo) GetWeekendDayOff() WeekendDayOff { return w.WeekendDayOff }
func (w *WorkInfo) GetSuperiorID() string           { return w.SuperiorID }

func (w *WorkInfo) Validate() *rest_err.RestErr {
	if w.Team == "" {
		return rest_err.NewBadRequestError("Team is required")
	}
	if w.Position == "" {
		return rest_err.NewBadRequestError("Position is required")
	}
	if w.SuperiorID == "" {
		return rest_err.NewBadRequestError("Superior ID is required")
	}
	return nil
}
