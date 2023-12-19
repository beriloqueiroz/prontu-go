package domain

import "time"

type Session struct {
	Id              string
	PatientIds      []string
	ProfessionalIds []string
	StartDate       time.Time
	TimeInMinutes   time.Duration
	Amount          float64
	EndDate         time.Time
	Notes           string
	Cids            []Cid
	Forms           []Form
}

type Cid struct {
	Id          string
	Code        string
	Name        string
	Observation string
}

type Form struct {
	Id   string
	Name string
	Link string
}
