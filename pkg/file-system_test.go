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
This is the first sentence of the content.
This is the second sentence.

This is the second paragraph.`

		secondPost = `This is the title of the second post
2013-Mar-10
picture2.jpg
bird,fly
-----
This is the first sentence of the content.`
	)

	dirFS := fstest.MapFS{
		"first-post.md":  {Data: []byte(firstPost)},
		"second post.md": {Data: []byte(secondPost)},
	}

	posts, err := New(dirFS)
	is.NoErr(err)

	t.Run("it returns a post for each (valid) file", func(t *testing.T) {
		is.Equal(len(posts), len(dirFS))
	})

	t.Run("it parses the first post correctly", func(t *testing.T) {
		expectedContent := "<p>This is the first sentence of the content.\nThis is the second sentence.</p>\n\n<p>This is the second paragraph.</p>\n"

		if string(posts[0].Content) != expectedContent {
			t.Errorf("got %q, want %q",posts[0].Content, expectedContent)
		}
		is.Equal(posts[0].Tags, []string{"cat", "dog"})
		is.Equal(posts[0].Picture, "picture.jpg")
		is.Equal(posts[0].Title, "This is the title")
		is.Equal(posts[0].Date, time.Date(2013, 03, 03, 0, 0, 0, 0, time.UTC))
	})
}
