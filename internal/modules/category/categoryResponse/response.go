package categoryresponse

type CategoryResponse struct {
	ID           uint   `json:"id"`
	CategoryType string `json:"category"`
}

type CategoriesResponse struct {
	Data []CategoryResponse `json:"categories"`
}
