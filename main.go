package main

import (
	"net/http"

	"github.com/AhmadRafly23/go-product-crud/handler"
	"github.com/AhmadRafly23/go-product-crud/model"
	"github.com/AhmadRafly23/go-product-crud/repository"
	"github.com/AhmadRafly23/go-product-crud/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    ge := gin.New()

	ge.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			map[string]any{
				"status": "OK!",
			})
	})


	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// melakukan migrasi / DDL
	// membuat table user
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}

	productPgRepo := &repository.ProductPgRepo{DB: db}
	productService := &service.ProductService{ProductPgRepo: productPgRepo}
	productHandler := &handler.ProductHandler{ProductService: productService}

	apiV1 := ge.Group("/api/v1")
	productGroup := apiV1.Group("/products")
	productGroup.GET("", productHandler.Get)
	productGroup.POST("", productHandler.Create)

	if err := ge.Run(":8080"); err != nil {
		panic(err)
	}

}
