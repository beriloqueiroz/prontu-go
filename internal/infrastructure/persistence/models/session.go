package infrastructure

import (
	"time"

	"gorm.io/gorm"
)

type Origin string

const (
	Google Origin = "google"
	None   Origin = "none"
)

type Session struct {
	gorm.Model
	Id            string `gorm:"primaryKey"`
	Patients      []Patient
	Professionals []Professional
	StartDate     time.Time
	TimeInMinutes time.Duration
	Amount        float64
	EndDate       time.Time
	Notes         string
	Cids          string
	Forms         string
	Origin        Origin
	ExternalId    string
	Location      string
}

func (s *Session) Validate() (valid bool, msg string) {
	msg, valid = "", true
	// if s.StartDate.Before(time.Now()) {
	// 	msg += msg + "Data inicial inválida; "
	// 	valid = false
	// }
	// if !s.EndDate.IsZero() && s.EndDate.Before(time.Now()) {
	// 	msg += msg + "Data final inválida; "
	// 	valid = false
	// }
	if s.Amount < 0 {
		msg += msg + "Valor inválido; "
		valid = false
	}
	if len(s.Patients) == 0 {
		msg += msg + "Nenhum paciente selecionado; "
		valid = false
	}
	if len(s.Professionals) == 0 {
		msg += msg + "Nenhum profissional selecionado; "
		valid = false
	}
	return valid, msg
}

type Patient struct {
	gorm.Model
	Id        uint `gorm:"primaryKey;autoIncrement"`
	PatientId string
	SessionId string
}

type Professional struct {
	gorm.Model
	Id             uint `gorm:"primaryKey;autoIncrement"`
	ProfessionalId string
	SessionID      string
}

var Sessions []Session
