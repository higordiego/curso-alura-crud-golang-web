package main

import (
	"log"
	"net/http"
	"text/template"
)

// Produto - struct
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	db := db.conectaComBancoDeDados()
	selectDeTodosOsprodutos, err := db.Query("select * from produtos")
	if err != nil {
		log.Println(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for selectDeTodosOsprodutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = selectDeTodosOsprodutos.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			log.Println(err.Error())
		}
		p.Nome = nome
		p.Quantidade = quantidade
		p.ID = id
		p.Descricao = descricao
		p.Preco = preco
		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)

	defer db.Close()
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
