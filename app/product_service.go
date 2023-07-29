package app

type ProductService struct {
	ProductPersistence ProductPersistenceInterface
}

func (ps *ProductService) Get(id string) (ProductInterface, error) {
	product, err := ps.ProductPersistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()

	if err != nil {
		return &Product{}, err
	}
	result, err := ps.ProductPersistence.Save(product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ps *ProductService) Enable(p ProductInterface) (ProductInterface, error) {
	err := p.Enable()
	if err != nil {
		return &Product{}, err
	}
	result, err := ps.ProductPersistence.Save(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ps *ProductService) Disable(p ProductInterface) (ProductInterface, error) {
	err := p.Disable()
	if err != nil {
		return &Product{}, err
	}
	result, err := ps.ProductPersistence.Save(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}
