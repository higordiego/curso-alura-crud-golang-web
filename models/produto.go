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

// CriarNovoProduto - struct
func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2 , $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

// DeletarProduto - struct models
func DeletarProduto(idProduto string) {
	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err)
	}

	deletarOProduto.Exec(idProduto)

	defer db.Close()
}

// EditaProduto - struct
func EditaProduto(idProduto string) Produto {
	db := db.ConectaComBancoDeDados()
	
	produtoBanco, err := db.Query("select id, nome, descricao, preco, quantidade from produtos where id = $1", idProduto)

	if err != nil {
		panic(err)
	}
	
	produtoParaAtualizar := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.ID = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Quantidade = quantidade
		produtoParaAtualizar.Preco = preco
	}

	defer db.Close()

	return produtoParaAtualizar
}