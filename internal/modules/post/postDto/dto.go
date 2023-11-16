package postdto

import (
	postmodel "crud/internal/modules/post/postModel"
	postresponse "crud/internal/modules/post/postResponse"
	"crud/pkg/helper"
)

var category map[uint]string

func DtOPostResponse(request postmodel.PostModel) postresponse.PostResponse {
	return postresponse.PostResponse{
		Id:        request.Id,
		UserId:    request.UserID,
		Title:     request.Title,
		Body:      request.Body,
		CreatedAt: helper.Now(request.CreatedAt),
		Category:  request.CategoryType,
	}
}

func DtoPostsResponse(posts []postmodel.PostModel) postresponse.PostsResponse {
	var response postresponse.PostsResponse

	for _, post := range posts {
		response.Data = append(response.Data, DtOPostResponse(post))
	}

	return response
}

func DtOCreatePost(posts postmodel.Posts) postresponse.CreatePostResponse {
	return postresponse.CreatePostResponse{
		UserId:    posts.UserID,
		Title:     posts.Title,
		Body:      posts.Body,
		CreatedAt: helper.Now(posts.CreatedAt),
		Category:  posts.CategoryID,
	}
}

func DtOUpdatePostResponse(request postmodel.Posts) postresponse.UpdatePostResponse {
	return postresponse.UpdatePostResponse{
		ID:         int(request.ID),
		UserId:     request.UserID,
		Title:      request.Title,
		Body:       request.Body,
		CategoryId: request.CategoryID,
	}
}

func DtODeletePostResponse(request postmodel.Posts) postresponse.DeletedPostResponse {
	return postresponse.DeletedPostResponse{
		ID:         int(request.ID),
		UserId:     request.UserID,
		Title:      request.Title,
		Body:       request.Body,
		CategoryId: request.CategoryID,
	}
}
