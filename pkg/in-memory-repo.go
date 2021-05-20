package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"personal-blog/pkg/blog"
	"sort"
)

type InMemoryRepository struct {
	posts []blog.Post
}

//TODO: test this. Look at new go testing docs. fishy that you have hard coded file paths

func NewInMemoryRepository() (*InMemoryRepository, error) {
	blogFiles, err := ioutil.ReadDir("posts")
	if err != nil {
		return nil, fmt.Errorf("cannot read the posts directory: %s", err)
	}

	var posts []blog.Post
	for _, file := range blogFiles {
		newPost, err := createPostFromFile(file.Name())
		if err != nil {
			return nil, fmt.Errorf("cannot create a new post: %s", err)
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

func createPostFromFile(file string) (blog.Post, error) {
	f, err := os.Open(fmt.Sprintf("../../cmd/web/posts/%s", file))
	if err != nil {
		return blog.Post{}, fmt.Errorf("cannot open the file %s: %s", file, err)
	}

	defer f.Close()

	return blog.CreatePost(f)
}
