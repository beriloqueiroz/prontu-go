package infrastructure

import "time"

type SessionControllerDto struct {
	Id              string        `json:"id"`
	PatientIds      []string      `json:"patientIds"`
	ProfessionalIds []string      `json:"ProfessionalIds"`
	StartDate       time.Time     `json:"startDate"`
	TimeInMinutes   time.Duration `json:"timeInMinutes"`
	Amount          float64       `json:"amount"`
	EndDate         time.Time     `json:"endDate"`
	Notes           string        `json:"notes"`
	Cids            []string      `json:"cids"`
	Forms           []string      `json:"forms"`
}

type DefaultError struct {
	Title string `json:"title"`
}
