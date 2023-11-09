package postresponse

type PostResponse struct {
	UserId        string `json:"user_id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	CreatedAt     string `json:"created_at"`
	Category_Type string `json:"category"`
}

type PostsResponse struct {
	Data []PostResponse
}

type UpdatePostResponse struct {
	ID            int    `json:"id"`
	UserId        string `json:"user_id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	Category_Type string `json:"category"`
}

type DeletedPostResponse struct {
	ID            int    `json:"id"`
	UserId        string `json:"user_id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	Category_Type string `json:"category"`
}
