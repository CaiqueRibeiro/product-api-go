package database

import (
	"github.com/CaiqueRibeiro/product-api/internal/entity"
)

type ProductRepositoryInterface interface {
	Create(product *entity.Product) error
	FindByID(id string) (*entity.Product, error)
	Delete(id string) error
	FindAll(page, limit int, sort string) ([]*entity.Product, error)
}
