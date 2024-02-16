package handlers

import (
	"net/http"
	"time"

	"github.com/CaiqueRibeiro/product-api/internal/dto"
	"github.com/CaiqueRibeiro/product-api/internal/entity"
	"github.com/CaiqueRibeiro/product-api/internal/infra/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandlers struct {
	Repository database.UserRepositoryInterface
}

func NewUserHandlers(repository database.UserRepositoryInterface) *UserHandlers {
	return &UserHandlers{
		Repository: repository,
	}
}

func (h *UserHandlers) CreateUser(c *gin.Context) {
	var user dto.CreateUserInput
	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Repository.Create(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *UserHandlers) GetJWT(c *gin.Context) {
	jwtSecret := c.MustGet("jwt").(string)
	jwtExpiresIn := c.MustGet("jwtExpiresIn").(int)

	var u dto.GetJWTInput
	if c.Bind(&u) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.Repository.FindByEmail(u.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	isValid := user.ValidatePassword(u.Password)
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	accessToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
