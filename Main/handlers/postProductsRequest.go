package handler

import (
	"net/http"

	model "Main/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func PostProductsRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.RequestPostProducts
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product := model.Products{
			Title:      request.Title,
			Price:      request.Price,
			Stock:      request.Stock,
			CategoryID: request.CategoryID,
		}

		validate := validator.New()
		if err := validate.Struct(product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		responseProduct := model.PostProductsRequest{
			Id:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:  product.CreatedAt,
		}

		c.JSON(http.StatusCreated, responseProduct)
	}
}
