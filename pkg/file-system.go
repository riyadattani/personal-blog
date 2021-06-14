package pkg

import (
	"fmt"
	"io/fs"
	"personal-blog/pkg/blog"
	"personal-blog/pkg/event"
	"sort"
)

//TODO: what is the best way to refactor this

func NewPosts(postsDir fs.FS) (blog.Posts, error) {
	dir, err := fs.ReadDir(postsDir, ".")
	if err != nil {
		return nil, fmt.Errorf("cannot read the posts directory: %s", err)
	}
	return getSortedPosts(postsDir, dir)
}

func getSortedPosts(postsDir fs.FS, dir []fs.DirEntry) ([]blog.Post, error) {
	var posts []blog.Post
	for _, file := range dir {
		post, err := newPostFromFile(postsDir, file.Name())
		if err != nil {
			return nil, fmt.Errorf("cannot create a new post: %s", err)
		}
		posts = append(posts, post)
	}

	return sortByDate(posts), nil
}

func newPostFromFile(postsDir fs.FS, fileName string) (blog.Post, error) {
	f, err := postsDir.Open(fileName)
	defer f.Close()

	if err != nil {
		return blog.Post{}, fmt.Errorf("cannot open the file %s: %s", fileName, err)
	}

	return blog.NewPost(f)
}

func sortByDate(posts []blog.Post) []blog.Post {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	return posts
}

func NewEvents(eventsDir fs.FS) ([]event.Event, error) {
	dir, err := fs.ReadDir(eventsDir, ".")
	if err != nil {
		return nil, fmt.Errorf("cannot read the events directory: %s", err)
	}
	return getSortedEvents(eventsDir, dir)
}

func getSortedEvents(eventsDir fs.FS, dir []fs.DirEntry) ([]event.Event, error) {
	var events []event.Event
	for _, file := range dir {
		event, err := newEventFromFile(eventsDir, file.Name())
		if err != nil {
			return nil, fmt.Errorf("cannot create a new event: %s", err)
		}
		events = append(events, event)
	}

	return sortEventsByDate(events), nil
}

func sortEventsByDate(events []event.Event) []event.Event {
	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.After(events[j].Date)
	})
	return events
}

func newEventFromFile(eventsDir fs.FS, fileName string) (event.Event, error) {
	f, err := eventsDir.Open(fileName)
	defer f.Close()

	if err != nil {
		return event.Event{}, fmt.Errorf("cannot open the file %s: %s", fileName, err)
	}

	return event.New(f)
}
