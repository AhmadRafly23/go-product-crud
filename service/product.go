package service

import (
	"github.com/AhmadRafly23/go-product-crud/model"
	"github.com/AhmadRafly23/go-product-crud/repository"
)

type ProductService struct {
	ProductPgRepo *repository.ProductPgRepo
}

func (u *ProductService) Get() ([]*model.Product, error) {
	return u.ProductPgRepo.Get()
}

func (u *ProductService) Create(product *model.Product) error {
	return u.ProductPgRepo.Create(product)
}