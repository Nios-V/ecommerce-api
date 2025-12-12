package products

import "time"

type Category struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Active bool   `gorm:"default:true" json:"active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Products []Product `gorm:"many2many:category_products;" json:"products,omitempty"`
}

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	Active      bool    `gorm:"default:true" json:"active"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"default:0" json:"stock"`

	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `json:"category"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
