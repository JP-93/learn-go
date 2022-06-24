package module

import (
	"fmt"

	dbc "alura/db"
)

type Product struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Qtd             int
}

func ConnectDB() {
	db := dbc.ConectionDB()
	defer db.Close()
}

func Getproduto() []Product {

	db := dbc.ConectionDB()
	defer db.Close()

	selectAll, err := db.Query("select * from produto")
	if err != nil {
		fmt.Errorf("Error on select %s", err)
	}

	p := Product{}
	produtos := []Product{}

	for selectAll.Next() {
		var id, qtd int
		var nome, descricao string
		var preco float64

		err = selectAll.Scan(&id, &nome, &descricao, &preco, &qtd)
		if err != nil {
			fmt.Errorf("Error on table %s", err)
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Qtd = qtd
		produtos = append(produtos, p)
	}
	return produtos
}

func CreateNewProduct(nome, descricao string, preco float64, qtd int) {
	db := dbc.ConectionDB()
	defer db.Close()

	insertData, err := db.Prepare("INSERT INTO produto(nome, descricao, preco, qtd)values ($1, $2, $3,$4)")
	if err != nil {
		fmt.Errorf("Error on insert %s", err)
	}
	insertData.Exec(nome, descricao, preco, qtd)
}

func DeleteProduct(id string) {
	db := dbc.ConectionDB()
	defer db.Close()

	deleteproduct, err := db.Prepare("DELETE FROM produto WHERE id=$1")
	if err != nil {
		fmt.Errorf("Error delete %s", err)
	}

	deleteproduct.Exec(id)
}
