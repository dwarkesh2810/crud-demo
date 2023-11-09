package postservice

import (
	postdto "crud/internal/modules/post/postDto"
	postmodel "crud/internal/modules/post/postModel"
	postrepository "crud/internal/modules/post/postRepository"
	postrequest "crud/internal/modules/post/postRequest"
	postresponse "crud/internal/modules/post/postResponse"
	"crud/utils"
)

type PostService struct {
	postRepo postrepository.PostRepositoryInterface
}

func New() *PostService {
	return &PostService{
		postRepo: postrepository.New(),
	}
}

func (postService *PostService) GetPosts() postresponse.PostsResponse {
	posts := postService.postRepo.List()
	return postdto.DtoPostsResponse(posts)
}

func (postService *PostService) StorePost(userId string, request postrequest.PostCreateRequest) (postresponse.PostResponse, error) {
	var post postmodel.Posts
	var response postresponse.PostResponse

	post.UserId = userId
	post.Title = request.Title
	post.Body = request.Body
	post.CategoryType = request.Category
	post.CreatedAt = utils.Now()

	newPost, err := postService.postRepo.Create(post)

	if err != nil {
		return response, err
	}

	return postdto.DtOPostResponse(newPost), nil

}
