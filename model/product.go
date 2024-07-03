package model

import "time"

type (
	Product struct {
		ID        uint64     `json:"id" gorm:"column:id;autoIncrement"`
		Name      string     `json:"name" gorm:"column:name"`
		Price     int 		`json:"price" gorm:"column:price"`
		DeletedAt *time.Time `json:"-" gorm:"-"`
	}

	ProductCreate struct {
		Name      string     `json:"name"`
		Price     int 		`json:"price"`
	}

	ProductUpdate struct {
		Name      string     `json:"name"`
		Price     int 		`json:"price"`
	}
)
