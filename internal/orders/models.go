package orders

import (
	"time"

	"github.com/Nios-V/ecommerce/api/internal/identity"
	"github.com/Nios-V/ecommerce/api/internal/products"
)

type Order struct {
	ID     uint        `gorm:"primaryKey" json:"id"`
	UserID uint        `gorm:"not null" json:"user_id"`
	Total  float64     `gorm:"not null" json:"total"`
	Status OrderStatus `gorm:"type:order_status_type;default:'PENDING'" json:"status"`

	ShippingAddressID uint             `gorm:"not null" json:"shipping_address_id"`
	ShippingAddress   identity.Address `json:"shipping_address"`

	BillingAddressID uint             `gorm:"not null" json:"billing_address_id"`
	BillingAddress   identity.Address `json:"billing_address"`

	Items   []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	Payment Payment     `gorm:"foreignKey:OrderID" json:"payment"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	OrderID         uint    `gorm:"not null" json:"order_id"`
	ProductID       uint    `gorm:"not null" json:"product_id"`
	Quantity        int     `gorm:"default:1" json:"quantity"`
	PriceAtPurchase float64 `gorm:"not null" json:"price"`

	Product products.Product `json:"product"`
}

type Payment struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	OrderID   uint          `gorm:"not null" json:"order_id"`
	Amount    float64       `gorm:"not null" json:"amount"`
	Method    PaymentMethod `gorm:"type:payment_method_type;not null" json:"method"`
	Status    PaymentStatus `gorm:"type:payment_status_type;default:'PENDING'" json:"status"`
	PaidAt    time.Time     `json:"paid_at"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
