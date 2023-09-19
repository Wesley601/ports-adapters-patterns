package product_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesley601/ports-adapters-patterns/internal/product"
)

func TestProduct_Enable(t *testing.T) {
	productImpl := product.NewProduct("1", "Product 1", product.DISABLE, 1000)
	err := productImpl.Enable()
	assert.Nil(t, err)
	assert.Equal(t, product.ENABLE, productImpl.GetStatus())
}
