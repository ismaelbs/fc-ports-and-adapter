package app

import (
	"errors"

	govalidator "github.com/asaskevich/govalidator"
	"github.com/dgryski/trifles/uuid"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Enable(p ProductInterface) (ProductInterface, error)
	Disable(p ProductInterface) (ProductInterface, error)
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
}

type ProductWriter interface {
	Save(p ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductWriter
	ProductReader
}

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Status string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
}

func NewProduct() *Product {
	product := Product{
		ID:     uuid.UUIDv4(),
		Status: DISABLED,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {

	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("Status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("Price must be greater than zero")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("Price must be greater than zero to enabled the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("Price must be zero to disabled the product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
