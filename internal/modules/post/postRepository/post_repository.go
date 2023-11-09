package postrepository

import (
	postmodel "crud/internal/modules/post/postModel"
	"crud/pkg/database"
	"errors"
	"log"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func New() *PostRepository {
	return &PostRepository{
		DB: database.Connection(),
	}
}

func (postRepository *PostRepository) List() []postmodel.Posts {
	var posts []postmodel.Posts
	postRepository.DB.Order("created_at desc").Find(&posts)
	return posts
}

func (postRepository *PostRepository) Create(post postmodel.Posts) (postmodel.Posts, error) {
	var newPost postmodel.Posts
	result := postRepository.DB.Create(&post).Scan(&newPost)
	if result.Error != nil {
		return newPost, errors.New("unexpected Database Error")
	}

	log.Print("new post", newPost)
	return newPost, nil
}

func (postRepository *PostRepository) Update(post postmodel.Posts) (postmodel.Posts, error) {
	return postmodel.Posts{}, nil
}
