package blog

type InMemoryRepository struct {
	blogs []Blog
}

func NewInMemoryRepository() *InMemoryRepository {
	var blogs []Blog
	blogs = append(blogs, NewBlog("This is my first Blog"))
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

