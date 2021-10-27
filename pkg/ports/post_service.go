package ports

import "personal-blog/pkg/blog"

type PostService interface {
	GetPosts() []blog.Post
	GetPost(title string) (blog.Post, error)
}
