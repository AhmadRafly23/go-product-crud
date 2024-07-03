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

func (u *ProductService) Update(id uint64, product *model.ProductUpdate) error {
	return u.ProductPgRepo.Update(id, product)
}

func (u *ProductService) Delete(id uint64) error {
	return u.ProductPgRepo.Delete(id)
}