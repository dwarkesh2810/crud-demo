package categoryrequest

type CategoryRequest struct {
	Category string `json:"category" db:"category_type"`
}

type CategoryDeleteRequest struct {
	Category string `json:"category" db:"category_type"`
}
