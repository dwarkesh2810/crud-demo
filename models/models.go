package models

type Posts struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	UserId       string `json:"user_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Body         string `json:"body" gorm:"not null"`
	CategoryType string `json:"category" gorm:"not null"`
	CreatedAt    string `json:"created_at" gorm:"not null"`
}

type Users struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    string `json:"user_id" gorm:"not null"`
	FirstName string `json:"first_name" gorm:"not null" form:"first_name"`
	LastName  string `json:"last_name" gorm:"not null" form:"last_name"`
	Email     string `json:"email" gorm:"not null;uniqueIndex" form:"email"`
	Password  string `json:"password" gorm:"not null" form:"password"`
	CreatedAt string `json:"created_at" gorm:"not null"`
	UserRole  string `json:"user_type" gorm:"not null"`
}

type Category struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Category_Type string `json:"category" gorm:"not null" db:"category_type"`
}

type Response struct {
	Success int         `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
