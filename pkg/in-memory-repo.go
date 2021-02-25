package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"personal-blog/pkg/blog"
	"sort"
)

type InMemoryRepository struct {
	posts []blog.Post
}

func NewInMemoryRepository() *InMemoryRepository {
	blogFiles, err := ioutil.ReadDir("posts")
	if err != nil {
		log.Fatal(fmt.Sprint("Cannot read the posts directory -  ", err))
	}

	var posts []blog.Post
	for _, post := range blogFiles {
		posts = append(posts, blog.NewPost(post.Name()))
	}
	return &InMemoryRepository{posts: posts}
}


func (i *InMemoryRepository) GetPost(title string) blog.Post {
	for _, post := range i.posts {
		if post.Title == title {
			return post
		}
	}
	return blog.NewPost("This does not exist")
}


func (i *InMemoryRepository) GetPosts() []blog.Post {
	posts := i.posts
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts
}

