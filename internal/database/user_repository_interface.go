package database

import (
	"github.com/CaiqueRibeiro/product-api/internal/entity"
)

type UserRepositoryInterface interface {
	Create(product *entity.User) error
	FindByID(id string) (*entity.User, error)
	Delete(id string) error
	FindAll(page, limit int, sort string) ([]*entity.User, error)
}
