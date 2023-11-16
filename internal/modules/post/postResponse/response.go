package postresponse

type PostResponse struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	Category  string `json:"category"`
}

type CreatePostResponse struct {
	UserId    uint   `json:"user_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	Category  uint   `json:"category"`
}

type PostsResponse struct {
	Data []PostResponse `json:"posts"`
}

type UpdatePostResponse struct {
	ID         int    `json:"id"`
	UserId     uint   `json:"user_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryId uint   `json:"category"`
}

type DeletedPostResponse struct {
	ID         int    `json:"id"`
	UserId     uint   `json:"user_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryId uint   `json:"category"`
}
