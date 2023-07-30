package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ismaelbs/fc-ports-and-adapter/adapters/cli"
	mock_app "github.com/ismaelbs/fc-ports-and-adapter/app/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 10.0
	productId := "123"
	productStatus := "enabled"

	productMock := mock_app.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_app.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf("Product created with id %s", productMock.GetID())
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product enabled with id %s", productMock.GetID())
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product disabled with id %s", productMock.GetID())
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product ID: %s\n Name: %s \n Price: %f\n Status: %s\n", productMock.GetID(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus())
	result, err = cli.Run(service, "", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
