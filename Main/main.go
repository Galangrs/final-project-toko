package main

import (
	config "Main/configs"
	handler "Main/handlers"
	middelware "Main/middelwares"
	model "Main/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	configDB := config.GetDBConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", configDB.Host, configDB.Port, configDB.User, configDB.DBName, configDB.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&model.Users{}, &model.Categories{}, &model.TransactionHistories{}, &model.Products{}); err != nil {
		log.Fatal("Failed to auto migrate:", err)
	}

	r.POST("/users/register", handler.PostRegisterRequest(db))

	r.POST("/users/login", handler.PostLoginRequest(db))

	r.Use(middelware.JWTMiddleware())

	r.POST("/users/topup", handler.PostTopupRequest(db))

	r.POST("/categories", handler.PostCategoriesRequest(db))

	r.GET("/categories", handler.GetCategoriesRequest(db))

	r.PATCH("/categories/:id", handler.PatchCategoriesRequest(db))

	r.DELETE("/categories/:id", handler.DeleteCategoriesRequest(db))

	r.POST("/products", middelware.Authorize(db), handler.PostProductsRequest(db))

	r.GET("/products", middelware.Authorize(db), handler.GetProductsRequest(db))

	r.PUT("/products/:id", middelware.Authorize(db), handler.PutProductsRequest(db))

	r.DELETE("/products/:id", middelware.Authorize(db), handler.DeleteProductsRequest(db))

	r.POST("/transactions", middelware.Authorize(db), handler.PostTransactionsRequest(db))

	r.GET("/transactions/my-transaction", middelware.Authorize(db), handler.GetTransactionsRequest(db))

	r.GET("/transactions/user-transactions", middelware.Authorize(db), handler.GetTransactionsAdminRequest(db))

	r.Run(":8080")
}
