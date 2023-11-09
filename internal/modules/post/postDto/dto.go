package postdto

import (
	postmodel "crud/internal/modules/post/postModel"
	postresponse "crud/internal/modules/post/postResponse"
)

func DtOPostResponse(request postmodel.Posts) postresponse.PostResponse {
	return postresponse.PostResponse{
		UserId:        request.UserId,
		Title:         request.Title,
		Body:          request.Body,
		CreatedAt:     request.CreatedAt,
		Category_Type: request.CategoryType,
	}
}

func DtoPostsResponse(posts []postmodel.Posts) postresponse.PostsResponse {
	var response postresponse.PostsResponse

	for _, post := range posts {
		response.Data = append(response.Data, DtOPostResponse(post))
	}

	return response
}


