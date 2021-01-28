package blog

import (
	"fmt"
	"io/ioutil"
	"log"
)

type InMemoryRepository struct {
	blogs []Blog
}

func NewInMemoryRepository() *InMemoryRepository {
	blogFiles, err := ioutil.ReadDir("blogs")
	if err != nil {
		log.Fatal(fmt.Sprint("Cannot read the blogs directory -  ", err))
	}

	var blogs []Blog
	for _, blog := range blogFiles {
		blogs = append(blogs, NewBlog(blog.Name()))
	}
	return &InMemoryRepository{blogs: blogs}
}


func (i *InMemoryRepository) GetBlog(id string) Blog {
	for _, blog := range i.blogs {
		if blog.ID==id {
			return blog
		}
	}
	return NewBlog("This does not exist")
}


func (i *InMemoryRepository) GetBlogs() []Blog {
	return i.blogs
}

