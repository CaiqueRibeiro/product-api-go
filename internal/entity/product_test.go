package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductEntity(t *testing.T) {
	t.Run("should create a new Product", func(t *testing.T) {
		createdAt := time.Now()
		product, err := NewProduct("Product 1", 10.5, createdAt)

		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, product.Name, "Product 1")
		assert.Equal(t, product.Price, 10.5)
		assert.Equal(t, product.CreatedAt, createdAt)
	})

	t.Run("return error if product was created with invalid value", func(t *testing.T) {
		createdAt := time.Now()
		product, err := NewProduct("Product 1", 0, createdAt)

		assert.NotNil(t, err)
		assert.Nil(t, product)
	})

	t.Run("Return error if name is not informed", func(t *testing.T) {
		product, err := NewProduct("", 10.5, time.Now())

		assert.NotNil(t, err)
		assert.EqualError(t, err, "name is required")
		assert.Nil(t, product)
	})

	t.Run("Return error if price is not informed", func(t *testing.T) {
		product, err := NewProduct("Product 1", 0.0, time.Now())

		assert.NotNil(t, err)
		assert.EqualError(t, err, "price is required")
		assert.Nil(t, product)
	})

	t.Run("Return error if price is negative", func(t *testing.T) {
		product, err := NewProduct("Product 1", -10.0, time.Now())

		assert.NotNil(t, err)
		assert.EqualError(t, err, "price is invalid")
		assert.Nil(t, product)
	})
}
