package pkg

import (
	"github.com/matryer/is"
	"testing"
	"testing/fstest"
	"time"
)

func TestBlog(t *testing.T) {
	is := is.New(t)

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

	posts, err := New(dirFS)
	is.NoErr(err)

	t.Run("it returns a post for each (valid) file", func(t *testing.T) {
		is.Equal(len(posts), len(dirFS))
	})

	t.Run("it parses the latest post correctly", func(t *testing.T) {
		latestPost := posts[0]

		is.Equal(latestPost.Title, "This is the title")
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
}
