package blog_test

import (
	"github.com/matryer/is"
	"personal-blog/pkg/blog"
	"strings"
	"testing"
	"time"
)

func TestBlog(t *testing.T) {
	is := is.New(t)
	t.Run("it should split the markdown file into the metadata and the content", func(t *testing.T) {
		markdownDoc := `This is the title
2013-Mar-03
picture.jpg
cat,dog
-----
This is the first sentence of the content.
This is the second sentence.

This is the second paragraph.`

		post, _ := blog.NewPost(strings.NewReader(markdownDoc))

		expectedBody := "<p>This is the first sentence of the content.\nThis is the second sentence.</p>\n\n<p>This is the second paragraph.</p>\n"
		expectedDate := time.Date(2013, 03, 03, 0, 0, 0, 0, time.UTC)

		is.Equal(string(post.Content), expectedBody)
		is.Equal(post.Title, "This is the title")
		is.Equal(post.Date, expectedDate)
		is.Equal(post.Picture, "picture.jpg")
		is.Equal(post.Tags, []string{"cat", "dog"})
		is.Equal(post.URLTitle, "This-is-the-title")
	})

	//TODO: test error case
}
