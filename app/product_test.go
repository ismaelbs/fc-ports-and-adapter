package app_test

import (
	"testing"

	"github.com/dgryski/trifles/uuid"
	"github.com/ismaelbs/fc-ports-and-adapter/app"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enabled(t *testing.T) {
	product := app.Product{}
	product.Name = "product"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "Price must be greater than zero to enabled the product", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := app.Product{}
	product.Name = "product"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	err = product.Disable()

	require.Equal(t, "Price must be zero to disabled the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.Name = "product"
	product.Status = app.DISABLED
	product.Price = 10
	product.ID = uuid.UUIDv4()

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "Status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	product.Price = -1
	_, err = product.IsValid()
	require.Equal(t, "Price must be greater than zero", err.Error())
}
