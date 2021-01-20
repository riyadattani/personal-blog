package blog

import "github.com/google/uuid"

type Blog struct {
	ID string
	Name string
}

func NewBlog(name string) Blog {
	return Blog{
		ID:      uuid.New().String(),
		Name:    name,
	}
}