package model

import "time"

type (
	Product struct {
		ID        uint64     `json:"id" gorm:"column:id;autoIncrement"`
		UserId	  uint64	 `json:"id_user" gorm:"column:user_id"`
		Name      string     `json:"name" gorm:"column:name"`
		Price     int 		 `json:"price" gorm:"column:price"`
		DeletedAt *time.Time `json:"-" gorm:"-"`
	}

	ProductCreate struct {
		Name      string    `json:"name"`
		Price     int 		`json:"price"`
		UserId	  uint64	`json:"user_id"`
	}

	ProductUpdate struct {
		Name      string    `json:"name"`
		Price     int 		`json:"price"`
	}
)
