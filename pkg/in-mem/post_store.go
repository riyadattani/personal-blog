package in_mem

import (
	"errors"
	"io/fs"
	"personal-blog/pkg/blog"
	"personal-blog/pkg/file_system"
)

type PostService interface {
	GetPosts() []blog.Post
	GetPost(title string) (blog.Post, error)
}

type PostSore struct {
	posts []blog.Post
}

func NewPostStore(postsDir fs.FS) (*PostSore, error) {
	posts, err := file_system.NewPosts(postsDir)
	if err != nil {
		return nil, err
	}

	return &PostSore{posts: posts}, nil
}

func (i *PostSore) GetPost(urlTitle string) (blog.Post, error) {
	for _, post := range i.posts {
		if post.URLTitle == urlTitle {
			return post, nil
		}
	}

	return blog.Post{}, errors.New("blog not found")
}

func (i *PostSore) GetPosts() []blog.Post {
	return i.posts
}
