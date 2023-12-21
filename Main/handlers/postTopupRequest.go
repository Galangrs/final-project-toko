package handler

import (
	"fmt"
	"net/http"

	helper "Main/helpers"
	model "Main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostTopupRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("ID")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
			return
		}
		userIDUint, ok := userID.(uint)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		var request model.RequestPostTopup
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user model.Users
		db.Where("id =?", userIDUint).First(&user)
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}
		user.Balance += request.Balance
		db.Save(&user)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Your account has been successfully updated to %s", helper.FormatRupiah(user.Balance))})
	}
}
