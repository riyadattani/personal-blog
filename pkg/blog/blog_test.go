package blog_test

import (
	"personal-blog/pkg/blog"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestBlog(t *testing.T) {
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
		expectedTitle := `This is the title`
		expectedDate := time.Date(2013, 03, 03, 0, 0, 0, 0, time.UTC)
		expectedPic := `picture.jpg`
		expectedTags := []string{"cat", "dog"}

		if string(post.Content) != expectedBody {
			t.Errorf("got %q, want %q", post.Content, expectedBody)
		}

		if post.Title != expectedTitle {
			t.Errorf("got %q, want %q", post.Title, expectedTitle)
		}

		if post.Date != expectedDate {
			t.Errorf("got %q, want %q", post.Date, expectedDate)
		}

		if post.Picture != expectedPic {
			t.Errorf("got %q, want %q", post.Picture, expectedPic)
		}

		if !reflect.DeepEqual(post.Tags, expectedTags) {
			t.Errorf("got %q, want %q", post.Tags, expectedTags)
		}
	})

	//TODO: test error case
}
