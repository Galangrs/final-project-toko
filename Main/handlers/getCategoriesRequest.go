package handler

import (
	"fmt"
	"net/http"

	model "Main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCategoriesHandler handles the GET request to retrieve categories and their products.
func GetCategoriesRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []model.Categories

		if err := db.Find(&categories).Error; err != nil {
			// Log the error for debugging
			fmt.Println("Error retrieving categories:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
			return
		}

		responseCategories := make([]model.GetCategoriesRequest, len(categories))
		for i, category := range categories {
			responseCategory := model.GetCategoriesRequest{
				Id:          category.ID,
				Type:        category.Type,
				SoldProduct: category.SoldProduct,
				CreatedAt:   category.CreatedAt,
				UpdatedAt:   category.UpdatedAt,
			}

			// Load products for each category
			if err := db.Where("category_id = ?", category.ID).Find(&category.Products).Error; err != nil {
				fmt.Printf("Error loading products for category %d: %s\n", category.ID, err)
				continue
			}

			if len(category.Products) > 0 {
				responseCategory.Products = make([]model.ProductInCategoryRequest, len(category.Products))
				for j, product := range category.Products {
					responseProduct := model.ProductInCategoryRequest{
						Id:        product.ID,
						Title:     product.Title,
						Price:     product.Price,
						Stock:     product.Stock,
						CreatedAt: product.CreatedAt,
					}
					responseCategory.Products[j] = responseProduct
				}
			}
			responseCategories[i] = responseCategory
		}

		c.JSON(http.StatusOK, responseCategories)
	}
}
