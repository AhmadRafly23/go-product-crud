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

func (s *ProductPgRepo) Update(id uint64, productUpdate *model.ProductUpdate) error {
	err := s.DB.Debug().
		Where("id = ?", id).
		Updates(&model.Product{
			Name:   productUpdate.Name,
			Price:    productUpdate.Price,
		}).Error
	return err
}

func (s *ProductPgRepo) Delete(id uint64) error {
	err := s.DB.Debug().
		Where("id = ?", id).
		Delete(&model.Product{}).Error
	return err
}