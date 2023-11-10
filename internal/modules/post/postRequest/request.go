package postrequest

type PostCreateRequest struct {
	Title    string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Body     string `form:"body" json:"body" binding:"required,min=3,max=10000"`
	Category string `form:"category" json:"category" binding:"required,min=3,max=20" db:"category_type"`
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

type UserAndCategoryRequest struct {
	UserID   string `json:"user_id"`
	Category string `json:"category" db:"category_type"`
}
