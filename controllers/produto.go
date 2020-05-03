package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/higordiego/curso-alura-crud-golang-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// TodosOsProdutos - handler
func TodosOsProdutos(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscarTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}

// New - handler
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

// Insert - handler
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão na quantidade: ", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
		http.Redirect(w, r, "/", 301)
	}
}
