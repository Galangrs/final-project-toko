package handler

import (
	"net/http"

	model "Main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductsRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []model.Products
		err := db.Find(&products).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		responseProducts := []model.GetProductsRequest{}

		for _, product := range products {
			responseProduct := model.GetProductsRequest{
				Id:         product.ID,
				Title:      product.Title,
				Price:      product.Price,
				Stock:      product.Stock,
				CategoryID: product.CategoryID,
				CreatedAt:  product.CreatedAt,
			}
			responseProducts = append(responseProducts, responseProduct)
		}
		c.JSON(200, responseProducts)
	}
}
