package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductEntity(t *testing.T) {
	tests := []struct {
		description string
		productName string
		price       float64
		createdAt   time.Time
		wantErr     bool
		errMessage  string
	}{
		{
			description: "should create a new Product",
			productName: "Product 1",
			price:       10.5,
			createdAt:   time.Now(),
			wantErr:     false,
		},
		{
			description: "return error if product was created with invalid value",
			productName: "Product 1",
			price:       0,
			createdAt:   time.Now(),
			wantErr:     true,
			errMessage:  "price is required",
		},
		{
			description: "Return error if name is not informed",
			productName: "",
			price:       10.5,
			createdAt:   time.Now(),
			wantErr:     true,
			errMessage:  "name is required",
		},
		{
			description: "Return error if price is not informed",
			productName: "Product 1",
			price:       0.0,
			createdAt:   time.Now(),
			wantErr:     true,
			errMessage:  "price is required",
		},
		{
			description: "Return error if price is negative",
			productName: "Product 1",
			price:       -10.0,
			createdAt:   time.Now(),
			wantErr:     true,
			errMessage:  "price is invalid",
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.description, func(t *testing.T) {
			product, err := NewProduct(scenario.productName, scenario.price, scenario.createdAt)

			if scenario.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, scenario.errMessage)
				assert.Nil(t, product)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, product.Name, scenario.productName)
				assert.Equal(t, product.Price, scenario.price)
				assert.Equal(t, product.CreatedAt, scenario.createdAt)
			}
		})
	}
}
