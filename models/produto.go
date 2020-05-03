package models

import (
	"log"

	"github.com/higordiego/curso-alura-crud-golang-web/db"
)

// Produto - struct
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// BuscarTodosOsProdutos - struct
func BuscarTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
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
	defer db.Close()
	return produtos
}
