package app_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ismaelbs/fc-ports-and-adapter/app"
	mock_app "github.com/ismaelbs/fc-ports-and-adapter/app/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil)

	service := app.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Get("1")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Create("product 1", 10.0)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Enabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_app.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)

	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := app.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Disabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_app.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := app.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
