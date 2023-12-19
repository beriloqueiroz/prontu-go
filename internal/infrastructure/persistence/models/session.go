package infrastructure

import "time"

type Session struct {
	Id             string        `json:"id"`
	PatientIds     []string      `json:"patientIds"`
	ProfessionalId []string      `json:"professionalId"`
	StartDate      time.Time     `json:"startDate"`
	TimeInMinutes  time.Duration `json:"timeInMinutes"`
	Amount         float64       `json:"amount"`
	EndDate        time.Time     `json:"endDate"`
	Notes          string        `json:"notes"`
	Cids           []Cid         `json:"cids"`
	Forms          []Form        `json:"forms"`
}

type Cid struct {
	Id          string `json:"id"`
	PatientId   string `json:"patientId"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Observation string `json:"observation"`
}

type Form struct {
	Id        string `json:"id"`
	PatientId string `json:"patientId"`
	Name      string `json:"name"`
	Link      string `json:"link"`
}

var Sessions []Session
