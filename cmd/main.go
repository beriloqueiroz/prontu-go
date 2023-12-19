package main

import (
	"fmt"
	"time"

	models "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/models"

	routes "github.com/beriloqueiroz/prontu-go/internal/infrastructure/api/routes"
)

func main() {
	models.Sessions = []models.Session{
		{
			Id:             "1",
			PatientIds:     []string{"112244", "223355"},
			ProfessionalId: []string{"1111", "2353"},
			StartDate:      time.Now(),
			TimeInMinutes:  10,
			Amount:         100.5,
			EndDate:        time.Now(),
			Notes:          "ta dnalsnda",
			Cids: []models.Cid{
				{Code: "5544", Name: "doença eeede", Observation: "não sei"},
			},
			Forms: []models.Form{
				{Name: "Anamnese", Link: "https://teste.com.be"},
			},
		},
		{
			Id:             "2",
			PatientIds:     []string{"111"},
			ProfessionalId: []string{"1111", "2353"},
			StartDate:      time.Now(),
			TimeInMinutes:  101,
			Amount:         150.5,
			EndDate:        time.Now(),
			Notes:          "ta kbhkb",
			Cids: []models.Cid{
				{Code: "5544", Name: "doença eeede", Observation: "não sei"},
			},
			Forms: []models.Form{
				{Name: "Anamnese", Link: "https://teste.com.be"},
			},
		},
	}
	fmt.Println("Iniciando o servidor")
	routes.HandleRequest()
}
