package postmodel

type Posts struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	UserId       string `json:"user_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Body         string `json:"body" gorm:"not null"`
	CategoryType string `json:"category" gorm:"not null"`
	CreatedAt    string `json:"created_at" gorm:"not null"`
}

// type PostsWithID struct {
// 	ID         uint   `json:"id" gorm:"primaryKey"`
// 	UserId     string `json:"user_id" gorm:"not null"`
// 	Title      string `json:"title" gorm:"not null"`
// 	Body       string `json:"body" gorm:"not null"`
// 	CategoryID int    `json:"category_id" gorm:"not null"`
// 	CreatedAt  string `json:"created_at" gorm:"not null"`
// }

// type Posts struct {
// 	ID         uint   `json:"id" gorm:"primaryKey"`
// 	UserId     string `json:"user_id" gorm:"not null;index;references:UserId"`
// 	Title      string `json:"title" gorm:"not null"`
// 	Body       string `json:"body" gorm:"not null"`
// 	CategoryID uint   `json:"category" gorm:"not null"`
// 	CreatedAt  string `json:"created_at" gorm:"not null"`

// 	User     usermodel.Users
// 	Category categorymodel.Category `gorm:"foreignKey:CategoryID"`
// }
