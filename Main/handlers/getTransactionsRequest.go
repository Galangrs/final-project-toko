package handler

import (
	"fmt"
	"net/http"

	model "Main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTransactionsRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("ID")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in the context"})
			return
		}

		var user model.Users
		resultUser := db.Where("id = ?", userID).First(&user)
		if resultUser.Error != nil {
			if resultUser.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			return
		}

		var transactions []model.TransactionHistories
		if err := db.Where("user_id = ?", user.ID).Find(&transactions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions: " + err.Error()})
			return
		}

		responseTransactions := make([]model.GetTransactionsRequestResponse, len(transactions))
		for i, transaction := range transactions {
			responseTransaction := model.GetTransactionsRequestResponse{
				Id:         transaction.ID,
				ProductID:  transaction.ProductID,
				UserID:     transaction.UserID,
				Quantity:   transaction.Quantity,
				TotalPrice: transaction.TotalPrice,
			}

			if err := db.Model(&transaction).Association("Product").Find(&transaction.Product); err != nil {
				fmt.Printf("Error loading product details for transaction %d: %s\n", transaction.ID, err)
				continue
			}

			if transaction.Product.ID != 0 {
				responseTransaction.Product = model.ProductInTransactionRequest{
					Id:        transaction.Product.ID,
					Title:     transaction.Product.Title,
					Price:     transaction.Product.Price,
					Stock:     transaction.Product.Stock,
					CreatedAt: transaction.Product.CreatedAt,
					UpdatedAt: transaction.Product.UpdatedAt,
				}
			}

			responseTransactions[i] = responseTransaction
		}

		c.JSON(http.StatusOK, responseTransactions)
	}
}
