package main

import (
	"fmt"
	"time"

	models "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/models"

	routes "github.com/beriloqueiroz/prontu-go/internal/infrastructure/api/routes"
	migration "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/migration"
)

func main() {
	models.Sessions = []models.Session{
		{
			Id:            "1",
			Patients:      []string{"2", "1"},
			Professionals: []string{"123"},
			StartDate:     time.Now(),
			TimeInMinutes: 10,
			Amount:        100.5,
			EndDate:       time.Now(),
			Notes:         "ta dnalsnda",
			Cids: []models.Cid{
				{Code: "5544", Name: "doença eeede", Observation: "não sei"},
			},
			Forms: "1jabskdjas.ndaslk",
		},
		{
			Id:            "2",
			Patients:      []string{"2", "41"},
			Professionals: []string{"1423"},
			StartDate:     time.Now(),
			TimeInMinutes: 101,
			Amount:        150.5,
			EndDate:       time.Now(),
			Notes:         "ta kbhkb",
			Cids: []models.Cid{
				{Code: "5a544", Name: "doença eeedeaa", Observation: "nãao sei"},
			},
			Forms: "1jabskdjas.ndaslaaaaak",
		},
	}
	fmt.Println("Iniciando o servidor")
	migration.InitMigration()
	routes.HandleRequest()
}
