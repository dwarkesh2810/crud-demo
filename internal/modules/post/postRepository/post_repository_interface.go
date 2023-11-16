package postrepository

import (
	postmodel "crud/internal/modules/post/postModel"
)

type PostRepositoryInterface interface {
	List() []postmodel.PostModel 
	Create(postmodel.Posts) (postmodel.Posts, error)
	Update(postmodel.Posts) (postmodel.Posts, error)
	GetPostById(uint) (postmodel.Posts, error)
	Delete(post postmodel.Posts) (postmodel.Posts, error)
}
