package cart

import (
	"time"

	"github.com/Nios-V/ecommerce/api/internal/products"
)

type Cart struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"not null" json:"user_id"`

	Items []CartItem `gorm:"foreignKey:CartID" json:"items"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CartID    uint    `gorm:"not null" json:"cart_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"default:1" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`

	Product products.Product `json:"product"`
}
