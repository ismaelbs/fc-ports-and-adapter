package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/ismaelbs/fc-ports-and-adapter/app"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db}
}

func (p *ProductDb) Get(id string) (app.ProductInterface, error) {
	var product app.Product
	stmt, err := p.db.Prepare("SELECT id, name, status, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Status, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
