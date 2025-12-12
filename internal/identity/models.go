package identity

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Active    bool   `gorm:"default:true" json:"active"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `json:"-"`

	Roles     []Role    `gorm:"many2many:user_roles;" json:"roles"`
	Addresses []Address `gorm:"foreignKey:UserID" json:"addresses"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Active bool   `gorm:"default:true" json:"active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Address struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	Active  bool   `gorm:"default:true" json:"active"`
	Street  string `gorm:"size:255;not null" json:"street"`
	City    string `gorm:"size:100;not null" json:"city"`
	State   string `gorm:"size:100;not null" json:"state"`
	ZipCode string `gorm:"size:20" json:"zip_code"`
	Country string `gorm:"size:100;not null" json:"country"`

	IsDefaultShipping bool `gorm:"default:false" json:"is_default_shipping"`
	IsDefaultBilling  bool `gorm:"default:false" json:"is_default_billing"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
