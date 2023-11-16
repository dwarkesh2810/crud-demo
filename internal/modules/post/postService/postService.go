package postservice

import (
	postdto "crud/internal/modules/post/postDto"
	postmodel "crud/internal/modules/post/postModel"
	postrepository "crud/internal/modules/post/postRepository"
	postrequest "crud/internal/modules/post/postRequest"
	postresponse "crud/internal/modules/post/postResponse"
	"errors"
	"time"
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

func (postService *PostService) CreatePost(userId uint, request postrequest.PostCreateRequest) (postresponse.CreatePostResponse, error) {
	var post postmodel.Posts
	var response postresponse.CreatePostResponse

	post.UserID = userId
	post.Title = request.Title
	post.Body = request.Body
	post.CategoryID = uint(request.Category)
	post.CreatedAt = time.Now().Unix()

	newPost, err := postService.postRepo.Create(post)

	if err != nil {
		return response, err
	}

	return postdto.DtOCreatePost(newPost), nil

}

func (postservice *PostService) UpdatePost(userId uint, request postrequest.PostUpdateRequest) (postresponse.UpdatePostResponse, error) {

	var response postresponse.UpdatePostResponse
	post, err := postservice.postRepo.GetPostById(request.ID)

	if err != nil {
		return response, err
	}

	if post.UserID != userId {
		return response, errors.New("you can not modify others post")
	}

	post.UserID = userId
	post.ID = request.ID
	post.Body = request.Body
	post.CategoryID = request.Category
	post.Title = request.Title

	updatedPost, err := postservice.postRepo.Update(post)
	if err != nil {
		return response, err
	}

	return postdto.DtOUpdatePostResponse(updatedPost), nil
}

func (postservice *PostService) DeletePost(userId uint, request postrequest.PostDeleteRequest) (postresponse.DeletedPostResponse, error) {

	var response postresponse.DeletedPostResponse
	post, err := postservice.postRepo.GetPostById(request.ID)
	if err != nil {
		return response, err
	}

	if post.UserID != userId {
		return response, errors.New("you can not modify others post")
	}

	deletePost, err := postservice.postRepo.Delete(post)
	if err != nil {
		return response, err
	}

	return postdto.DtODeletePostResponse(deletePost), nil
}
