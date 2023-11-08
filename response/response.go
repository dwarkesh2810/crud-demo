package response

type PostResponse struct {
	UserId        string `json:"user_id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	CreatedAt     string `json:"created_at"`
	Category_Type string `json:"category"`
}

type UserDataResponse struct {
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}