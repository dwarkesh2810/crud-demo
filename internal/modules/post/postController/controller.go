package postcontroller

import postservice "crud/internal/modules/post/postService"

type Controllers struct {
	postService postservice.PostServiceInterface
}

func New() *Controllers {
	return &Controllers{
		postService: postservice.New(),
	}
}
