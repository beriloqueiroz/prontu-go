package infrastructure

import (
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
	db.AutoMigrate(&models.Session{}, &models.Cid{}) //todo ver como tirar auto e rodar pelo .sql
}
