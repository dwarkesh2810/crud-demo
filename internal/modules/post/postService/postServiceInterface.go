package postservice

import (
	postrequest "crud/internal/modules/post/postRequest"
	postresponse "crud/internal/modules/post/postResponse"
)

type PostServiceInterface interface {
	GetPosts() postresponse.PostsResponse
	StorePost(string, postrequest.PostCreateRequest) (postresponse.PostResponse, error)
}
