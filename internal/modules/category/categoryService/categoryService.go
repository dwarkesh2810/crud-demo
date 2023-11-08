package categoryservice

import (
	postrepository "crud/internal/modules/post/postRepository"
)

type PostService struct {
	postRepo postrepository.PostRepositoryInterface
}

func New() *PostService {
	return &PostService{
		postRepo: postrepository.New(),
	}
}
