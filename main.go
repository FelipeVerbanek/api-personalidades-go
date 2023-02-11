package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Jose", Historia: "Tatata"},
		{Id: 2, Nome: "Joao", Historia: "Tatata"},
	}
	fmt.Println("Conectando ao banco de dados!")
	database.ConectaComBancoDados()

	fmt.Println("Iniciando servidor!")
	routes.HandleRequest()
}
