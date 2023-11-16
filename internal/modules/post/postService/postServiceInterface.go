package postservice

import (
	postrequest "crud/internal/modules/post/postRequest"
	postresponse "crud/internal/modules/post/postResponse"
)

type PostServiceInterface interface {
	GetPosts() postresponse.PostsResponse
	CreatePost(uint, postrequest.PostCreateRequest) (postresponse.CreatePostResponse, error)
	UpdatePost(uint, postrequest.PostUpdateRequest) (postresponse.UpdatePostResponse, error)
	DeletePost(uint, postrequest.PostDeleteRequest) (postresponse.DeletedPostResponse, error)
}
