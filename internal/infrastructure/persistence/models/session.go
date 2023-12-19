package infrastructure

import (
	"time"

	"gorm.io/gorm"
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
