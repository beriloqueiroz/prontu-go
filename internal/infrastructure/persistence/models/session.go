package infrastructure

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Id            string        `gorm:"primaryKey"`
	Patients      Patients      `gorm:"type:VARCHAR(255)"`
	Professionals Professionals `gorm:"type:VARCHAR(255)"`
	StartDate     time.Time
	TimeInMinutes time.Duration
	Amount        float64
	EndDate       time.Time
	Notes         string
	Cids          []Cid `gorm:"foreignKey:SessionId"`
	Forms         string
}

type Patients []string
type Professionals []string

type Cid struct {
	gorm.Model
	Id          uint `gorm:"primaryKey;autoIncrement";`
	PatientId   string
	Code        string
	Name        string
	Observation string
	SessionId   string
}

var Sessions []Session
