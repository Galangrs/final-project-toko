package middelware

import (
	"net/http"

	model "Main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Authorize(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("ID")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
			return
		}
		var user model.Users
		resultUser := db.Where("id = ?", userID).First(&user)
		if resultUser.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/password supplied"})
			return
		}

		if user.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
