package categorymodel

type Category struct {
	CategoryID   uint   `json:"id" gorm:"primaryKey"`
	CategoryType string `json:"category" gorm:"not null;uniqueIndex"`
}
