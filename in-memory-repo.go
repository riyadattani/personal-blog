package blog

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type InMemoryRepository struct {
	posts []Post
}

func NewInMemoryRepository() *InMemoryRepository {
	blogFiles, err := ioutil.ReadDir("posts")
	if err != nil {
		log.Fatal(fmt.Sprint("Cannot read the posts directory -  ", err))
	}

	var posts []Post
	for _, post := range blogFiles {
		posts = append(posts, NewPost(post.Name()))
	}
	return &InMemoryRepository{posts: posts}
}


func (i *InMemoryRepository) GetPost(title string) Post {
	for _, post := range i.posts {
		if post.Title == title {
			return post
		}
	}
	return NewPost("This does not exist")
}


func (i *InMemoryRepository) GetPosts() []Post {
	posts := i.posts
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts
}

