package postmodel

import (
	categorymodel "crud/internal/modules/category/categoryModel"
	usermodel "crud/internal/modules/user/userModel"
)

// type Posts struct {
// 	ID           uint   `json:"id" gorm:"primaryKey"`
// 	UserId       uint   `json:"user_id" gorm:"not null"`
// 	Title        string `json:"title" gorm:"not null"`
// 	Body         string `json:"body" gorm:"not null"`
// 	CategoryType string `json:"category" gorm:"not null"`
// 	CreatedAt    int64  `json:"created_at" gorm:"not null"`
// }

type Posts struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserID     uint   `json:"user_id" gorm:"not null;index;references:UserID"` // Added references tag
	Title      string `json:"title" gorm:"not null"`
	Body       string `json:"body" gorm:"not null"`
	CategoryID uint   `json:"category" gorm:"not null;index"`
	CreatedAt  int64  `json:"created_at" gorm:"not null"`

	User     usermodel.Users        `gorm:"foreignKey:UserID"`
	Category categorymodel.Category `gorm:"foreignKey:CategoryID"`
}

type PostModel struct {
	Id           uint   `json:"id"`
	UserID       uint   `json:"user_id"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	CategoryType string `json:"category"`
	CreatedAt    int64  `json:"created_at"`
}
