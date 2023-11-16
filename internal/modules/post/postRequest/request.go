package postrequest

type PostCreateRequest struct {
	Title    string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Body     string `form:"body" json:"body" binding:"required,min=3,max=10000"`
	Category uint   `form:"category" json:"category" binding:"required,numeric"`
}

type PostUpdateRequest struct {
	ID       uint   `form:"id" binding:"required,numeric"`
	Title    string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Body     string `form:"body" json:"body" binding:"required,min=3,max=10000"`
	Category uint   `form:"category" json:"category" binding:"required,numeric"`
}

type PostDeleteRequest struct {
	ID uint `form:"id" json:"id" binding:"required,numeric"`
}

type UserAndCategoryRequest struct {
	UserID   uint `json:"user_id"`
	Category uint `json:"category,numeric"`
}
