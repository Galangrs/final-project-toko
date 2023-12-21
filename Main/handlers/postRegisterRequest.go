package handler

import (
	"net/http"

	model "Main/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func PostRegisterRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.RequestPostRegister
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		user := model.Users{
			FullName: request.FullName,
			Email:    request.Email,
			Password: request.Password,
		}
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.Password = string(hashedPassword)
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register user"})
			return
		}
		responseUser := model.PostRegisterRequest{
			Id:        user.ID,
			Fullname:  user.FullName,
			Email:     user.Email,
			Password:  user.Password,
			Balance:   user.Balance,
			CreatedAt: user.CreatedAt,
		}
		c.JSON(http.StatusCreated, responseUser)
	}
}
