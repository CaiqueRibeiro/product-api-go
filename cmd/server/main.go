package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CaiqueRibeiro/product-api/configs"
	"github.com/CaiqueRibeiro/product-api/internal/entity"
	"github.com/CaiqueRibeiro/product-api/internal/infra/database"
	"github.com/CaiqueRibeiro/product-api/internal/infra/webserver/handlers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func Auth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Decode and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return // Abort() prevents to call the next middleware/handler, but the current one keeps going in next lines
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the expiry date
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
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

	userRepository := database.NewUserRepository(db)
	userHandler := handlers.NewUserHandlers(userRepository)

	productRepository := database.NewProductRepository(db)
	productHandler := handlers.NewProductHandlers(productRepository)

	r := gin.Default()

	r.Use(JWTConfigs(configs.JWTSecret, configs.JWTExpiresIn))

	r.POST("/users", userHandler.CreateUser)
	r.POST("/users/generate_token", userHandler.GetJWT)

	authorized := r.Group("/products")
	{
		authorized.Use(Auth(configs.JWTSecret))
		authorized.POST("/", productHandler.CreateProduct)
		authorized.GET("/", productHandler.FindAll)
	}

	r.Run(configs.WebServerPort)
}
