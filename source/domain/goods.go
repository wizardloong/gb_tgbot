package domain

import "time"

type Goods struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Price     float64   `gorm:"type:decimal(10,2);not null"`
	Data      string    `gorm:"type:text;not null"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
