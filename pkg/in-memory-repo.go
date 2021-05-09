package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"personal-blog/pkg/blog"
	"sort"
)

type InMemoryRepository struct {
	posts []blog.Post
}

func NewInMemoryRepository() (*InMemoryRepository, error) {
	blogFiles, err := ioutil.ReadDir("posts")
	if err != nil {
		return &InMemoryRepository{}, fmt.Errorf("cannot read the posts directory: %s", err)
	}

	var posts []blog.Post
	for _, post := range blogFiles {
		newPost, err := blog.NewPost(post.Name())
		if err != nil {
			return &InMemoryRepository{}, fmt.Errorf("cannot create a new post: %s", err)
		}
		posts = append(posts, newPost)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

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
