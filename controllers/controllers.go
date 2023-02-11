package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var newPersonalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&newPersonalidade)
	database.DB.Create(&newPersonalidade)
	json.NewEncoder(w).Encode(newPersonalidade)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	w.WriteHeader(204)
}
