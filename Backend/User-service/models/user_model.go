package model

import "time"


type User struct {
ID uint `gorm:"primaryKey" json:"id"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`
Name string `gorm:"size:100;not null" json:"name"`
Email string `gorm:"size:100;uniqueIndex;not null" json:"email"`
Password string `gorm:"size:255;not null" json:"-"`
}