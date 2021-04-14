package pkg

import (
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

func (i *InMemoryRepository) GetPost(title string) blog.Post {
	for _, post := range i.posts {
		if post.Title == title {
			return post
		}
	}
	//TODO handle the 404 - return an error (sentinal?) if err == RiyaNotFound, do something. What if someone god /blog/bob? - write a sad path test in server.test
	post, _ := blog.NewPost("This does not exist")
	return post
}

func (i *InMemoryRepository) GetPosts() []blog.Post {
	return i.posts
}
