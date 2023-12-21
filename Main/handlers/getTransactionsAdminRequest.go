package handler

import (
	"fmt"
	"net/http"

	model "Main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTransactionsAdminRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var transactions []model.TransactionHistories
		if err := db.Find(&transactions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions: " + err.Error()})
			return
		}

		responseTransactions := make([]model.GetTransactionsAdminRequestResponse, len(transactions))
		for i, transaction := range transactions {
			responseTransaction := model.GetTransactionsAdminRequestResponse{
				Id:         transaction.ID,
				ProductID:  transaction.ProductID,
				UserID:     transaction.UserID,
				Quantity:   transaction.Quantity,
				TotalPrice: transaction.TotalPrice,
			}

			if err := db.Preload("User").Preload("Product").Find(&transaction).Error; err != nil {
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

			if transaction.User.ID != 0 {
				responseTransaction.User = model.UserInTransactionRequest{
					Id:        transaction.User.ID,
					Email:     transaction.User.Email,
					FullName:  transaction.User.FullName,
					Balance:   transaction.User.Balance,
					CreatedAt: transaction.User.CreatedAt,
					UpdatedAt: transaction.User.UpdatedAt,
				}
			}
			responseTransactions[i] = responseTransaction
		}

		c.JSON(http.StatusOK, responseTransactions)
	}
}
