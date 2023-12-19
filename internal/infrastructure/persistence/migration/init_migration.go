package infrastructure

import (
	"time"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	models "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitMigration() {

	dsn := "host=localhost user=root password=root dbname=root port=5436 sslmode=disable TimeZone=America/Fortaleza"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Session{}, &models.Patient{}, &models.Professional{}) //todo ver como tirar auto e rodar pelo .sql
	seed(db)
}

func seed(db *gorm.DB) {
	Sessions := []models.Session{
		{
			Id: "1",
			Patients: []models.Patient{
				{PatientId: "123"},
			},
			Professionals: []models.Professional{
				{ProfessionalId: "123"},
				{ProfessionalId: "1234"},
			},
			StartDate:     time.Now(),
			TimeInMinutes: 10,
			Amount:        100.5,
			EndDate:       time.Now(),
			Notes:         "ta dnalsnda",
			Cids:          "Ass1, 123",
			Forms:         "1jabskdjas.ndaslk",
		},
		{
			Id: "2",
			Patients: []models.Patient{
				{PatientId: "123"},
			},
			Professionals: []models.Professional{
				{ProfessionalId: "123"},
				{ProfessionalId: "1234"},
			},
			StartDate:     time.Now(),
			TimeInMinutes: 101,
			Amount:        150.5,
			EndDate:       time.Now(),
			Notes:         "ta kbhkb",
			Cids:          "Ass1, 123",
			Forms:         "1jabskdjas.ndaslaaaaak",
		},
	}
	db.Create(Sessions)
}
