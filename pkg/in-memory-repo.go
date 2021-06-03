package pkg

import (
	"errors"
	"io/fs"
	"personal-blog/pkg/blog"
)

type InMemoryRepository struct {
	posts []blog.Post
}

func NewInMemoryRepository(postsDir fs.FS) (*InMemoryRepository, error) {
	posts, err := New(postsDir)
	if err != nil {
		return nil, err
	}

	return &InMemoryRepository{posts: posts}, nil
}

func (i *InMemoryRepository) GetPost(title string) (blog.Post, error) {
	for _, post := range i.posts {
		if post.Title == title {
			return post, nil
		}
	}

	return blog.Post{}, errors.New("blog not found")
}

func (i *InMemoryRepository) GetPosts() []blog.Post {
	return i.posts
}