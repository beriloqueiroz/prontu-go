package infrastructure

import (
	"os"
	"time"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	models "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	dsn := os.Getenv("DB_CONN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.Session{}, &models.Patient{}, &models.Professional{}) //todo ver como tirar auto e rodar pelo .sql
	// seed(DB)
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
			Origin:        "none",
			Location:      "",
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
			Location:      "",
			Origin:        "none",
		},
	}
	db.Create(Sessions)
}
