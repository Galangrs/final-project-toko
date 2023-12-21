package handler

import (
	"net/http"

	model "Main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteCategoriesRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryID := c.Param("id")

		var category model.Categories
		if err := db.First(&category, categoryID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}

		if err := db.Delete(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Category has been successfully deleted"})
	}
}
