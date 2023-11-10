package userresponse

type UserResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserResponse struct {
	UserResponse
	Token string `json:"token"`
}
