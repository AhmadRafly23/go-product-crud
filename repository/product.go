package repository

import (
	"github.com/AhmadRafly23/go-product-crud/model"
	"gorm.io/gorm"
)

type ProductPgRepo struct {
	DB *gorm.DB
}

func (s *ProductPgRepo) Get() ([]*model.Product, error) {
	products := []*model.Product{}
	err := s.DB.Debug().Find(&products).Error
	return products, err
}

func (s *ProductPgRepo) Create(product *model.Product) error {
	err := s.DB.Debug().Create(&product).Error
	return err
}

