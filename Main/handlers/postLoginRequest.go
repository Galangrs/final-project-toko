package handler

import (
	"net/http"

	helper "Main/helpers"
	model "Main/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func PostLoginRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.RequestPostLogin
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user model.Users
		result := db.Where("email = ?", request.Email).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/password supplied"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/password supplied"})
			return
		}

		roleString := string(user.Role)

		token, err := helper.GenerateJWTToken(user.ID, roleString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
