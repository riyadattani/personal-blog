package blog_test

import (
	"personal-blog/pkg/blog"
	"testing"
	"time"
)

func TestBlog(t *testing.T) {
	t.Run("it should return the title, date and content separately", func(t *testing.T) {
		markdownDoc := `About something else
2013-Mar-03
picture.jpg
cat,dog
-----
This is the about me thing`

		byteArray := []byte(markdownDoc)
		title, body, date, picture, _, _ := blog.CreatePost(byteArray)

		expectedBody := `This is the about me thing`
		expectedTitle := `About something else`
		expectedDate := `2013-Mar-03`
		expectedPic := `picture.jpg`
		//expectedTags := []string{"cat", "dog"}


		if string(body) != expectedBody {
			t.Errorf("got %q, want %q", body, expectedBody)
		}

		if title != expectedTitle {
			t.Errorf("got %q, want %q", title, expectedTitle)
		}

		if date != expectedDate {
			t.Errorf("got %q, want %q", date, expectedDate)
		}

		if picture != expectedPic {
			t.Errorf("got %q, want %q", picture, expectedPic)
		}
		//TODO: why doesnt this work
		//if tags != expectedTags {
		//	t.Errorf("got %q, want %q", tags, expectedTags)
		//}
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
