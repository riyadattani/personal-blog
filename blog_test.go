package blog

import (
	"strings"
	"testing"
)

//pass in a document, turn it to an io.Reader, call the function and it should return the meta data and the reader for the rest of the body.

func TestBlog(t *testing.T) {
	t.Run("it should return the meta data and the content separately", func(t *testing.T) {
		markdownDoc := `About
2006-01-15
-----
This is the about me thing`

		reader := strings.NewReader(markdownDoc)

		metaData := getMetaData(reader)

		expectedMetaData := `About
2006-01-15`

		got := metaData
		want := expectedMetaData

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}

