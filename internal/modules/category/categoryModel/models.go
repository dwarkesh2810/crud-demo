package categorymodel

type Category struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Category_Type string `json:"category" gorm:"not null;uniqueIndex" db:"category_type"`
}
