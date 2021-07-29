package pkg

import (
	"errors"
	"io/fs"
	"personal-blog/pkg/blog"
	"personal-blog/pkg/event"
	"personal-blog/pkg/file_system"
)

type InMemoryRepository struct {
	posts  []blog.Post
	events []event.Event
}

func NewInMemoryRepository(postsDir fs.FS, eventsDir fs.FS) (*InMemoryRepository, error) {
	posts, err := file_system.NewPosts(postsDir)
	if err != nil {
		return nil, err
	}
	events, err1 := file_system.NewEvents(eventsDir)
	if err1 != nil {
		return nil, err1
	}

	return &InMemoryRepository{posts: posts, events: events}, nil
}

func (i *InMemoryRepository) GetPost(urlTitle string) (blog.Post, error) {
	for _, post := range i.posts {
		if post.URLTitle == urlTitle {
			return post, nil
		}
	}

	return blog.Post{}, errors.New("blog not found")
}

func (i *InMemoryRepository) GetPosts() []blog.Post {
	return i.posts
}

func (i *InMemoryRepository) GetEvents() []event.Event {
	return i.events
}
