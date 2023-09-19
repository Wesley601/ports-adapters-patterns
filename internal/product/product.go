package product

import "errors"

const (
	DISABLE = "disabled"
	ENABLE  = "enabled"
)

type Product interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductImpl struct {
	id     string
	name   string
	status string
	price  float64
}

func NewProduct(id, name, status string, price float64) Product {
	return &ProductImpl{
		id:     id,
		name:   name,
		status: status,
		price:  price,
	}
}

func (p *ProductImpl) Enable() error {
	if p.price <= 0 {
		return errors.New("price must be greater than 0 to enable the product")
	}

	p.status = ENABLE
	return nil
}

func (p *ProductImpl) Disable() error {
	p.status = DISABLE
	return nil
}

func (p *ProductImpl) GetID() string {
	return p.id
}

func (p *ProductImpl) GetName() string {
	return p.name
}

func (p *ProductImpl) GetStatus() string {
	return p.status
}

func (p *ProductImpl) GetPrice() float64 {
	return p.price
}

func (p *ProductImpl) IsValid() (bool, error) {
	return true, nil
}
