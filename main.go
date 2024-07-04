package main

import (
	"net/http"

	"github.com/AhmadRafly23/go-product-crud/handler"
	"github.com/AhmadRafly23/go-product-crud/middleware"
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
	err = db.AutoMigrate(&model.User{}, &model.Product{})
	if err != nil {
		panic(err)
	}

	userPgRepo := &repository.UserPgRepo{DB: db}
	userService := &service.UserService{UserPgRepo: userPgRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	productPgRepo := &repository.ProductPgRepo{DB: db}
	productService := &service.ProductService{ProductPgRepo: productPgRepo}
	productHandler := &handler.ProductHandler{ProductService: productService}

	apiV1 := ge.Group("/api/v1")
	userGroup := apiV1.Group("/users")
	productGroup := apiV1.Group("/products")

	userGroup.POST("/register", userHandler.Create)
	userGroup.POST("/login", userHandler.Login)
	userGroup.Use(middleware.BearerAuthorization())
	userGroup.GET("", userHandler.Get)

	productGroup.Use(middleware.BearerAuthorization())
	productGroup.GET("", productHandler.Get)
	productGroup.POST("", productHandler.Create)
	productGroup.PUT("/:id", productHandler.Update)
	productGroup.DELETE("/:id", productHandler.Delete)

	ge.Run(":8080")

}
