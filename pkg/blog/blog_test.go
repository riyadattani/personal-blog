package blog_test

import (
	"personal-blog/pkg/blog"
	"strings"
	"testing"
	"time"
)

func TestBlog(t *testing.T) {
	t.Run("it should return the title, date and content separately", func(t *testing.T) {
		markdownDoc := `About
2013-Feb-03
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
2013-Mar-03
picture.jpg
-----
This is the about me thing`

		reader := strings.NewReader(markdownDoc)
		metaData := blog.GetMetaData(reader)

		title := metaData[0]
		expectedTitle := `About something else`

		date := metaData[1]
		expectedDate := `2013-Mar-03`

		picture := metaData[2]
		expectedPic := `picture.jpg`

		if title != expectedTitle {
			t.Errorf("got %q, want %q", title, expectedTitle)
		}

		if date != expectedDate {
			t.Errorf("got %q, want %q", date, expectedDate)
		}

		if picture != expectedPic {
			t.Errorf("got %q, want %q", picture, expectedPic)
		}
	})

	t.Run("Format the date correctly", func(t *testing.T) {
		date := "2013-Mar-03"
		const layout = "2006-Jan-02"
		formattedDate, _ := time.Parse(layout, date)
		expectedFormattedDate := time.Date(2013, 03, 03, 0, 0, 0, 0, time.UTC)

		if formattedDate != expectedFormattedDate {
			t.Errorf("got %q, want %q", formattedDate, expectedFormattedDate)
		}
	})
}
