package pkg

import (
	"github.com/matryer/is"
	"testing"
	"testing/fstest"
	"time"
)

func TestFileSystem(t *testing.T) {
	is := is.New(t)

	t.Run("articles", func(t *testing.T) {
		const (
			firstPost = `This is the title
2013-Mar-03
picture.jpg
cat,dog
-----
blah blah blah`

			secondPost = `This is the title of the second post
2013-Mar-01
picture2.jpg
bird,fly
-----
This is the first sentence of the content.`
		)

		dirFS := fstest.MapFS{
			"first-post.md":  {Data: []byte(firstPost)},
			"second-post.md": {Data: []byte(secondPost)},
		}

		posts, err := NewPosts(dirFS)
		is.NoErr(err)

		t.Run("it returns a post for each (valid) file", func(t *testing.T) {
			is.Equal(len(posts), len(dirFS))
		})

		t.Run("it parses the latest post correctly", func(t *testing.T) {
			latestPost := posts[0]

			is.Equal(latestPost.Title, "This is the title")
			is.Equal(latestPost.URLTitle, "This-is-the-title")
			is.Equal(latestPost.Date, time.Date(2013, 03, 03, 0, 0, 0, 0, time.UTC))
			is.Equal(latestPost.Picture, "picture.jpg")
			is.Equal(latestPost.Tags, []string{"cat", "dog"})
			is.Equal(string(latestPost.Content), "<p>blah blah blah</p>\n")
		})

		t.Run("parses the next post correctly", func(t *testing.T) {
			nextPost := posts[1]
			is.Equal(nextPost.Title, "This is the title of the second post")
			is.Equal(nextPost.Date, time.Date(2013, 03, 01, 0, 0, 0, 0, time.UTC))
		})
	})

	t.Run("Events", func(t *testing.T) {
		const (
			event1 = `This is the title
2013-Mar-03
www.google.com
cat,dog
-----
blah blah blah`

			event2 = `This is the title of the second post
2013-Mar-01
www.google.com
bird,fly
-----
This is the first sentence of the content.`
		)

		dirFS := fstest.MapFS{
			"event1.md":  {Data: []byte(event1)},
			"event2.md": {Data: []byte(event2)},
		}

		events, err := NewEvents(dirFS)
		is.NoErr(err)

		t.Run("it returns an event for each (valid) file", func(t *testing.T) {
			is.Equal(len(events), len(dirFS))
		})

		t.Run("it parses the latest event correctly", func(t *testing.T) {
			latestEvent := events[0]

			is.Equal(latestEvent.Title, "This is the title")
			is.Equal(latestEvent.Date, time.Date(2013, 03, 03, 0, 0, 0, 0, time.UTC))
			is.Equal(latestEvent.Link, "www.google.com")
			is.Equal(latestEvent.Tags, []string{"cat", "dog"})
			is.Equal(string(latestEvent.Description), "<p>blah blah blah</p>\n")
		})

		t.Run("parses the next event correctly", func(t *testing.T) {
			nextEvent := events[1]
			is.Equal(nextEvent.Title, "This is the title of the second post")
			is.Equal(nextEvent.Date, time.Date(2013, 03, 01, 0, 0, 0, 0, time.UTC))
		})

	})
}
