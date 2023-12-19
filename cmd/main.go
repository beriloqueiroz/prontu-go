package main

import (
	"fmt"

	routes "github.com/beriloqueiroz/prontu-go/internal/infrastructure/api/routes"
	migration "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/migration"
)

func main() {
	fmt.Println("Iniciando o servidor")
	migration.InitMigration()
	routes.HandleRequest()
}
