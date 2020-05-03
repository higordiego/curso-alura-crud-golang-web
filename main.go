package main

import (
	"net/http"
	"text/template"

	"github.com/higordiego/curso-alura-crud-golang-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscarTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
