package postservice

import (
	postrequest "crud/internal/modules/post/postRequest"
	postresponse "crud/internal/modules/post/postResponse"
)

type PostServiceInterface interface {
	GetPosts() postresponse.PostsResponse
	CreatePost(string, postrequest.PostCreateRequest) (postresponse.PostResponse, error)
	UpdatePost(string, postrequest.PostUpdateRequest) (postresponse.UpdatePostResponse, error)
	DeletePost(postrequest.PostDeleteRequest) (postresponse.DeletedPostResponse, error)
}
