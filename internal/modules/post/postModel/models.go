package postmodel

type Posts struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	UserId       string `json:"user_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Body         string `json:"body" gorm:"not null"`
	CategoryType string `json:"category" gorm:"not null"`
	CreatedAt    string `json:"created_at" gorm:"not null"`
}

