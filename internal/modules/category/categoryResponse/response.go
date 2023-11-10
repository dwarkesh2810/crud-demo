package categoryresponse

type CategoryResponse struct {
	ID       uint   `json:"id"`
	Category string `json:"category"`
}

type CategoriesResponse struct {
	Data []CategoryResponse `json:"data"`
}
