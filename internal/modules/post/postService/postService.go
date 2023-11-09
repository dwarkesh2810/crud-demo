package postservice

import (
	postdto "crud/internal/modules/post/postDto"
	postmodel "crud/internal/modules/post/postModel"
	postrepository "crud/internal/modules/post/postRepository"
	postrequest "crud/internal/modules/post/postRequest"
	postresponse "crud/internal/modules/post/postResponse"
	"crud/pkg/helper"
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

func (postService *PostService) CreatePost(userId string, request postrequest.PostCreateRequest) (postresponse.PostResponse, error) {
	var post postmodel.Posts
	var response postresponse.PostResponse

	post.UserId = userId
	post.Title = request.Title
	post.Body = request.Body
	post.CategoryType = request.Category
	post.CreatedAt = helper.Now()

	newPost, err := postService.postRepo.Create(post)

	if err != nil {
		return response, err
	}

	return postdto.DtOPostResponse(newPost), nil

}

func (postservice *PostService) UpdatePost(userId string, request postrequest.PostUpdateRequest) (postresponse.UpdatePostResponse, error) {

	var response postresponse.UpdatePostResponse
	post, err := postservice.postRepo.GetPostById(request.ID)
	post.UserId = userId
	post.ID = request.ID
	post.Body = request.Body
	post.CategoryType = request.Category
	post.Title = request.Title

	updatedPost, err := postservice.postRepo.Update(post)
	if err != nil {
		return response, err
	}

	return postdto.DtOUpdatePostResponse(updatedPost), nil
}

func (postservice *PostService) DeletePost(request postrequest.PostDeleteRequest) (postresponse.DeletedPostResponse, error) {

	var response postresponse.DeletedPostResponse
	post, err := postservice.postRepo.GetPostById(request.ID)
	if err != nil {
		return response, err
	}

	deletePost, err := postservice.postRepo.Delete(post)
	if err != nil {
		return response, err
	}

	return postdto.DtODeletePostResponse(deletePost), nil
}
