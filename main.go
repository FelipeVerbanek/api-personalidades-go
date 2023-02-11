package main

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Jose", Historia: "Tatata"},
		{Nome: "Joao", Historia: "Tatata"},
	}

	fmt.Println("Iniciando servidor!")
	routes.HandleRequest()
}
