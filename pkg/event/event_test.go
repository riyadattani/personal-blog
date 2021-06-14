package event

import (
	"github.com/matryer/is"
	"strings"
	"testing"
	"time"
)

func TestBlog(t *testing.T) {
	is := is.New(t)
	t.Run("Should create an event from a file", func(t *testing.T) {
		markdownDoc := `Event title
2013-Mar-03
https://www.riyadattani.com/
picture.jpg
cat,dog
-----
This is a short description of an event`

		event, _ := New(strings.NewReader(markdownDoc))

		expectedBody := "<p>This is a short description of an event</p>\n"
		is.Equal(string(event.Description), expectedBody)
		is.Equal(event.Title, "Event title")
		is.Equal(event.Link, "https://www.riyadattani.com/")
		is.Equal(event.Date, time.Date(2013, 03, 03, 0, 0, 0, 0, time.UTC))
		is.Equal(event.Picture, "picture.jpg")
		is.Equal(event.Tags, []string{"cat", "dog"})
	})

	//TODO: test error case
}