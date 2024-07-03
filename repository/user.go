package repository

import (
	"github.com/AhmadRafly23/go-product-crud/model"
	"gorm.io/gorm"
)

type UserPgRepo struct {
	DB *gorm.DB
}

func (s *UserPgRepo) Get() ([]*model.User, error) {
	users := []*model.User{}
	err := s.DB.Debug().Find(&users).Error
	return users, err
}

func (s *UserPgRepo) Create(student *model.User) error {
	err := s.DB.Debug().Create(&student).Error
	return err
}

func (s *UserPgRepo) Update(id uint64, studentUpdate *model.UserUpdate) error {
	err := s.DB.Debug().
		Where("id = ?", id).
		Updates(&model.User{
			Name:   studentUpdate.Name,
			Gender: studentUpdate.Gender,
		}).Error
	return err
}

func (s *UserPgRepo) Delete(id uint64) error {
	err := s.DB.Debug().
		Where("id = ?", id).
		Delete(&model.User{}).Error
	return err
}

func (s *UserPgRepo) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := s.DB.Debug().Where("email = ?", email).Find(&user).Error
	return user, err
}