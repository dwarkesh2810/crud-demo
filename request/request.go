package request

type PostCreateRequest struct {
	Title    string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Body     string `form:"body" json:"body" binding:"required,min=3,max=10000"`
	Category string `form:"category_type" json:"category_type" binding:"required,min=3,max=20" db:"category_type"`
}

type PostUpdateRequest struct {
	ID       uint   `form:"id" binding:"required,numeric"`
	Title    string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Body     string `form:"body" json:"body" binding:"required,min=3,max=10000"`
	Category string `form:"category" json:"category" binding:"required,min=3,max=20"`
}

type PostDeleteRequest struct {
	ID uint `form:"id" json:"id" binding:"required,numeric"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email,min=3,max=100"`
	Password string `form:"password" binding:"required,min=8"`
}

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=100"`
	LastName  string `json:"last_name" binding:"required,min=3,max=100" `
	Email     string `json:"email" binding:"required,email,min=3,max=100"`
	Password  string `json:"password" binding:"required,min=8"`
}

type CategoryCreateRequest struct {
	Category string `form:"category" db:"category_type"`
}

type CategoryDeleteRequest struct {
	Category string `form:"category_type" db:"category_type"`
}
