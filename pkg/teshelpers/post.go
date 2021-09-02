package teshelpers

import (
	"personal-blog/pkg/blog"
	"time"
)

type BlogPostBuilder struct {
	Post blog.Post
}

func NewPostBuilder() *BlogPostBuilder {
	return &BlogPostBuilder{Post: blog.Post{
		Title:    "This is the title of the builder blog",
		Content:  "This is the content of the builder blog",
		Date:     time.Now(),
		Picture:  "picture.png",
		Tags:     []string{"bob", "builder"},
		URLTitle: "This-is-the-title-of-the-builder-blog",
	},
	}
}

func (b *BlogPostBuilder) WithTitle(title string) *BlogPostBuilder {
	b.Post.Title = title
	return b
}

func (b *BlogPostBuilder) Build() blog.Post {
	return b.Post
}
