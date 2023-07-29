package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ismaelbs/fc-ports-and-adapter/adapters/db"
	"github.com/ismaelbs/fc-ports-and-adapter/app"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	insertProduct(Db)
}

func createTable(*sql.DB) {
	_, err := Db.Exec("CREATE TABLE IF NOT EXISTS products (id string, name string, status string, price REAL)")
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(*sql.DB) {
	_, err := Db.Exec("INSERT INTO products (id, name, status, price) VALUES ('abc', 'Product 1', 'disabled', 0)")
	if err != nil {
		log.Fatal(err)
	}
}

func TestProductdb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductdb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	product := app.NewProduct()
	product.Name = "Product 1"
	product.Price = 0

	productDb := db.NewProductDb(Db)
	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

	product.Status = "enabled"
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

}
