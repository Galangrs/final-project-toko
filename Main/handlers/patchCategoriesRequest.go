package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"

	model "Main/models"
)

func PatchCategoriesRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.RequestPatchCategories
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")

		var category model.Categories
		if err := db.First(&category, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		category.Type = request.Type

		validate := validator.New()
		if err := validate.Struct(category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Save(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
			return
		}

		respnseCategory := model.PatchCategoriesRequest{
			Id:          category.ID,
			Type:        category.Type,
			SoldProduct: category.SoldProduct,
			UpdatedAt:   category.UpdatedAt,
		}
		c.JSON(http.StatusOK, respnseCategory)
	}
}
