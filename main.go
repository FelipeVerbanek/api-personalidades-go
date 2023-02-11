package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {

	fmt.Println("Conectando ao banco de dados!")
	database.ConectaComBancoDados()

	fmt.Println("Iniciando servidor!")
	routes.HandleRequest()
}
