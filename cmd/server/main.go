package main

import (
	"github.com/CaiqueRibeiro/product-api/configs"
	"github.com/CaiqueRibeiro/product-api/internal/entity"
	"github.com/CaiqueRibeiro/product-api/internal/infra/database"
	"github.com/CaiqueRibeiro/product-api/internal/infra/webserver/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func JWTConfigs(jwtSecret string, jwtExpiresIn int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("jwt", jwtSecret)
		c.Set("jwtExpiresIn", jwtExpiresIn)
		c.Next()
	}
}

func main() {
	configs, err := configs.LoadConfig(".env")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	// productRepository := database.NewProductRepository(db)

	userRepository := database.NewUserRepository(db)
	userHandler := handlers.NewUserHandlers(userRepository)

	r := gin.New()

	r.Use(JWTConfigs(configs.JWTSecret, configs.JWTExpiresIn))

	r.POST("/users", userHandler.CreateUser)
	r.POST("/users/generate_token", userHandler.GetJWT)

	r.Run(configs.WebServerPort)
}
