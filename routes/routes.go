package routes

import (
	"net/http"

	"github.com/higordiego/curso-alura-crud-golang-web/controllers"
)

// CarregarRotas - handler loading routes
func CarregarRotas() {
	http.HandleFunc("/", controllers.TodosOsProdutos)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
