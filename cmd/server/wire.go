//go:build wireinject
// +build wireinject

package main

import (
	"github.com/CaiqueRibeiro/product-api/internal/infra/database"
	"github.com/CaiqueRibeiro/product-api/internal/infra/webserver/handlers"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var setProductRepositoryDepencency = wire.NewSet(
	database.NewProductRepository,
	wire.Bind(new(database.ProductRepositoryInterface), new(*database.ProductRepository)),
)

var setUserRepositoryDepencency = wire.NewSet(
	database.NewUserRepository,
	wire.Bind(new(database.UserRepositoryInterface), new(*database.UserRepository)),
)

func NewProductHandler(db *gorm.DB) *handlers.ProductHandler {
	wire.Build(
		setProductRepositoryDepencency,
		handlers.NewProductHandlers,
	)
	return &handlers.ProductHandler{}
}

func NewUserHandler(db *gorm.DB) *handlers.UserHandlers {
	wire.Build(
		setUserRepositoryDepencency,
		handlers.NewUserHandlers,
	)
	return &handlers.UserHandlers{}
}
