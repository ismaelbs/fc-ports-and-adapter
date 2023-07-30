package cli

import (
	"fmt"

	"github.com/ismaelbs/fc-ports-and-adapter/app"
)

func Run(service app.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	var result string
	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product created with id %s", product.GetID())
	case "enable":
		p, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err := service.Enable(p)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product enabled with id %s", product.GetID())
	case "disable":
		p, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err := service.Disable(p)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product disabled with id %s", product.GetID())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID: %s\n Name: %s \n Price: %f\n Status: %s\n", res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}
