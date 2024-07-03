package model

import "time"

type (
	GenderType string

	User struct {
		ID        uint64     `json:"id" gorm:"column:id;autoIncrement"`
		Password  string     `json:"-" gorm:"column:password"`
		Name      string     `json:"name" gorm:"column:name"`
		Email     string     `json:"email" gorm:"column:email"`
		Gender    GenderType `json:"gender" gorm:"column:gender"`
		DeletedAt *time.Time `json:"-" gorm:"-"`
	}

	UserUpdate struct {
		Name   string     `json:"name"`
		Gender GenderType `json:"gender"`
	}

	UserCreate struct {
		Name     string     `json:"name"`
		Password string     `json:"password"`
		Email    string     `json:"email"`
		Gender   GenderType `json:"gender"`
	}

	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)