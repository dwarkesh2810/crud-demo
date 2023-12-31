package usermodel

type Users struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserId    uint   `json:"user_id" gorm:"not null;uniqueIndex"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null;uniqueIndex"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"not null"`
	UserRole  string `json:"user_type" gorm:"not null"`
}
