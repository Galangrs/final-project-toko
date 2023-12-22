package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	model "Main/models"
)

func PostTransactionsRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("ID")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
			return
		}
		var request model.RequestPostTransactions
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tx := db.Begin()

		var product model.Products
		resultProduct := db.Where("id = ?", request.ProductID).First(&product)
		if resultProduct.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Product not found"})
			return
		}

		transaction := model.TransactionHistories{
			UserID:     int(userID.(uint)),
			ProductID:  request.ProductID,
			Quantity:   request.Quantity,
			TotalPrice: request.Quantity * product.Price,
		}

		if product.Stock < request.Quantity {
			tx.Rollback()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Insufficient stock"})
			return
		}
		if product.Stock == 0 {
			tx.Rollback()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "stock is 0"})
			return
		}
		product.Stock -= request.Quantity

		var user model.Users
		resultUser := db.Where("id = ?", userID).First(&user)
		if resultUser.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/password supplied"})
			return
		}
		if user.Balance < transaction.TotalPrice {
			tx.Rollback()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Insufficient balance"})
			return
		}
		user.Balance -= transaction.TotalPrice
		var category model.Categories
		if err := db.First(&category, product.ID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		category.SoldProduct = category.SoldProduct + request.Quantity
		if err := db.Save(&category).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
			return
		}
		responseTransaction := model.PostTransactionsRequest{
			TotalPrice:   transaction.TotalPrice,
			Quantity:     transaction.Quantity,
			ProductTitle: product.Title,
		}
		if err := db.Model(&user).Updates(user).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := db.Model(&product).Updates(product).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&transaction).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{
			"message":          "You have succesfully purchased the product",
			"transaction_bill": responseTransaction,
		})
	}
}
