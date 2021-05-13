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

		expectedBody :=  "<p>This is the first sentence of the content</p>\n"
		expectedTitle := `This is the title`
		expectedDate := time.Date(2013, 03, 03, 0, 0, 0, 0, time.UTC)
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

	//TODO: test error case
}
