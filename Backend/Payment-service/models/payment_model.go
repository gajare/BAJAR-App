package models

import "time"

type Payment struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID       uint64    `gorm:"not null" json:"order_id"`
	UserID        uint64    `gorm:"not null" json:"user_id"`
	Amount        float64   `gorm:"type:numeric(10,2)" json:"amount"`
	PaymentMethod string    `gorm:"size:50" json:"payment_method"`
	Status        string    `gorm:"size:50;default:'pending'" json:"status"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}
