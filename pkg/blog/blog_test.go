package blog_test

import (
	"personal-blog/pkg/blog"
	"strings"
	"testing"
)

func TestBlog(t *testing.T) {
	t.Run("it should return the title, date and content separately", func(t *testing.T) {
		markdownDoc := `About
2006-01-15
-----
This is the about me thing`

		byteArray := []byte(markdownDoc)
		contentBody := string(blog.GetContentBody(byteArray))
		expectedBody := `This is the about me thing`

		if contentBody != expectedBody {
			t.Errorf("got %q, want %q", expectedBody, contentBody)
		}
	})

	t.Run("it should read a line in the text", func(t *testing.T) {
		markdownDoc := `About something else
2006-01-15
-----
This is the about me thing`

		reader := strings.NewReader(markdownDoc)
		metaData := blog.GetMetaData(reader)

		title := metaData[0]
		expectedTitle := `About something else`

		date := metaData[1]
		expectedDate := `2006-01-15`

		if title != expectedTitle {
			t.Errorf("got %q, want %q", title, expectedTitle)
		}

		if date != expectedDate {
			t.Errorf("got %q, want %q", date, expectedDate)
		}
	})
}
