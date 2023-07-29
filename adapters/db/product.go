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

func (p *ProductDb) create(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO products (id, name, status, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetStatus(), product.GetPrice())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, status = ?, price = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetName(), product.GetStatus(), product.GetPrice(), product.GetID())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) Save(product app.ProductInterface) (app.ProductInterface, error) {
	var rows int
	p.db.QueryRow("Select count(*) from products where id=?", product.GetID()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}
