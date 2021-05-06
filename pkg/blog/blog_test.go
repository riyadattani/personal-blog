package blog_test

import (
	"personal-blog/pkg/blog"
	"reflect"
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
This is the first sentence of the content`

		byteArray := []byte(markdownDoc)
		metaData, body, _ := blog.CreatePost(byteArray)

		expectedBody := `This is the first sentence of the content`
		expectedTitle := `This is the title`
		expectedDate := `2013-Mar-03`
		expectedPic := `picture.jpg`
		expectedTags := []string{"cat", "dog"}

		if string(body) != expectedBody {
			t.Errorf("got %q, want %q", body, expectedBody)
		}

		if metaData.Title != expectedTitle {
			t.Errorf("got %q, want %q", metaData.Title, expectedTitle)
		}

		if metaData.Date != expectedDate {
			t.Errorf("got %q, want %q", metaData.Date, expectedDate)
		}

		if metaData.Picture != expectedPic {
			t.Errorf("got %q, want %q", metaData.Picture, expectedPic)
		}

		if !reflect.DeepEqual(metaData.Tags, expectedTags) {
			t.Errorf("got %q, want %q", metaData.Tags, expectedTags)
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
