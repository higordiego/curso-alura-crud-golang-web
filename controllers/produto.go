package controllers

import (
	"net/http"
	"text/template"

	"github.com/higordiego/curso-alura-crud-golang-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// TodosOsProdutos - handler
func TodosOsProdutos(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscarTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}
