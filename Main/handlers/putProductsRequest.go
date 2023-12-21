package handler

import (
	"net/http"

	helper "Main/helpers"
	model "Main/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func PutProductsRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var request model.RequestPutProducts
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newProduct := model.Products{
			Title:      request.Title,
			Price:      request.Price,
			Stock:      request.Stock,
			CategoryID: request.CategoryID,
		}

		var product model.Products
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		validate := validator.New()
		if err := validate.Struct(newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&product).Updates(newProduct).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		responseProduct := model.PutProductsRequest{
			Id:         product.ID,
			Title:      product.Title,
			Price:      helper.FormatRupiah(product.Price),
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		}
		c.JSON(http.StatusOK, gin.H{"product": responseProduct})
	}
}
