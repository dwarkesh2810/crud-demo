package postrequest

type PostCreateRequest struct {
	Title    string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Body     string `form:"body" json:"body" binding:"required,min=3,max=10000"`
	Category string `form:"category_type" json:"category_type" binding:"required,min=3,max=20" db:"category_type"`
}
