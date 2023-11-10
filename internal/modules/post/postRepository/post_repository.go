package postrepository

import (
	postmodel "crud/internal/modules/post/postModel"
	"crud/pkg/config"
	"errors"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func New() *PostRepository {
	return &PostRepository{
		DB: config.Connection(),
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

	return newPost, nil
}

func (postRepository *PostRepository) GetPostById(id uint) (postmodel.Posts, error) {
	var post postmodel.Posts
	result := config.DB.First(&post, "id = ?", id)

	if result.Error != nil {
		return post, errors.New("unexpected Database Error")
	}
	return post, nil
}

func (postRepository *PostRepository) Update(post postmodel.Posts) (postmodel.Posts, error) {

	postAfterUpdate := post

	result := config.DB.Save(&postAfterUpdate)

	if result.Error != nil {
		return postmodel.Posts{}, errors.New("Unexpected database error")
	}

	return postAfterUpdate, nil
}

func (postRepository *PostRepository) Delete(post postmodel.Posts) (postmodel.Posts, error) {

	postDeleted := post

	result := config.DB.Delete(&postDeleted)

	if result.Error != nil {
		return postmodel.Posts{}, errors.New("Unexpected database error")
	}

	return postDeleted, nil
}



// config.DB.Select("posts.id, posts.title, posts.body, categories.category_type, posts.created_at").Joins("join categories.id=posts.category_id").Where("category_id", id).Find(&posts).Scan(&results)