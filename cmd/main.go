package main

import (
	"fmt"

	routes "github.com/beriloqueiroz/prontu-go/internal/infrastructure/api/routes"
	database "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence"
)

func main() {
	fmt.Println("Iniciando o servidor")
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
