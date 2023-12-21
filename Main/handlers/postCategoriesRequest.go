package handler

import (
	"net/http"

	model "Main/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func PostCategoriesRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.RequestPostCategories
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		category := model.Categories{
			Type: request.Type,
		}
		validate := validator.New()
		if err := validate.Struct(category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
			return
		}
		responseCategory := model.PostCategoriesRequest{
			Id:          category.ID,
			Type:        category.Type,
			SoldProduct: category.SoldProduct,
			CreatedAt:   category.CreatedAt,
		}
		c.JSON(http.StatusCreated, responseCategory)
	}
}
