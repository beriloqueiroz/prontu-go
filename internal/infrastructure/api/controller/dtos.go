package infrastructure

import "time"

type SessionControllerDto struct {
	Id              string              `json:"id"`
	PatientIds      []string            `json:"patientIds"`
	ProfessionalIds []string            `json:"ProfessionalIds"`
	StartDate       time.Time           `json:"startDate"`
	TimeInMinutes   time.Duration       `json:"timeInMinutes"`
	Amount          float64             `json:"amount"`
	EndDate         time.Time           `json:"endDate"`
	Notes           string              `json:"notes"`
	Cids            []CidControllerDto  `json:"cids"`
	Forms           []FormControllerDto `json:"forms"`
}

type CidControllerDto struct {
	Id          string `json:"id"`
	PatientId   string `json:"patientId"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Observation string `json:"observation"`
}

type FormControllerDto struct {
	Id        string `json:"id"`
	PatientId string `json:"patientId"`
	Name      string `json:"name"`
	Link      string `json:"link"`
}
