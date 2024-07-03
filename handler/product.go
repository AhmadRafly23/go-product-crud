package handler

import (
	"net/http"
	"strconv"

	"github.com/AhmadRafly23/go-product-crud/model"
	"github.com/AhmadRafly23/go-product-crud/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func (u *ProductHandler) Get(ctx *gin.Context) {
	products, err := u.ProductService.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Message: "products fetched",
		Success: true,
		Data:    products,
	})
}

func (u *ProductHandler) Create(ctx *gin.Context) {
	// binding payload
	productCreate := model.ProductCreate{}
	if err := ctx.Bind(&productCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
	}
	// call service
	err := u.ProductService.Create(&model.Product{
		Name:  productCreate.Name,
		Price:   productCreate.Price,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	// response

	ctx.JSON(http.StatusCreated, model.Response{
		Message: "products created",
		Success: true,
	})
}

func (u *ProductHandler) Update(ctx *gin.Context) {
	// bind id from path param
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
	}
	id, _ := strconv.Atoi(idStr)
	// binding payload
	productUpdate := model.ProductUpdate{}
	if err := ctx.Bind(&productUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
	}
	// call service
	err := u.ProductService.Update(uint64(id), &productUpdate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	// response

	ctx.JSON(http.StatusCreated, model.Response{
		Message: "product updated",
		Success: true,
	})
}

func (u *ProductHandler) Delete(ctx *gin.Context){
	idStr := ctx.Param("id")

	if idStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
		return
	}

	id, _ := strconv.Atoi(idStr)

	err := u.ProductService.Delete(uint64(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.Response{
		Message: "success deleted",
		Success: true,
	})
}