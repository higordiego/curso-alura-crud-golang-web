package main

import (
	"net/http"

	"github.com/higordiego/curso-alura-crud-golang-web/routes"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
