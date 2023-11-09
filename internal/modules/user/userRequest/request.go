package userrequest

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=100"`
	LastName  string `json:"last_name" binding:"required,min=3,max=100" `
	Email     string `json:"email" binding:"required,email,min=3,max=100"`
	Password  string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email,min=3,max=100"`
	Password string `form:"password" binding:"required,min=8"`
}
