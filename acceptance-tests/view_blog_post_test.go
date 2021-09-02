package acceptance_tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"personal-blog/pkg"
	"personal-blog/pkg/blog"
	"personal-blog/pkg/http_api"
	"testing"
)
//TODO: use selinium to create the adapter - this is like cypress
type LocalBlogAdapter struct {
}

func (l LocalBlogAdapter) Publish(post blog.Post) error {
	//TODO: testing lib has a thing where it create a temp file and then deletes it when it is done
	file, err := os.Create("posts-test/acceptance-test.md")
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintf(file, `%s
2013-Mar-03
picture.jpg
cat,dog
-----
This is the first sentence of the content.
This is the second sentence.

This is the second paragraph.`, post.Title)
	return nil
}

func (l LocalBlogAdapter) ReadPost(urlTitle string) (bool, error) {
	config := pkg.NewConfig()
	config.PostsDir = "posts-test"
	config.EventsDir = "events-test"

	app, err := pkg.NewApplication(config)
	if err != nil {
		return false, err
	}
	server := http_api.NewServer(app.Config, &app.Handler)
	svr := httptest.NewServer(server.Handler)

	defer svr.Close()

	url := svr.URL + "/blog/" + urlTitle
	res, err := http.Get(url)
	if err != nil {
		return false, err
	}
	return res.StatusCode == http.StatusOK, nil
}

func TestViewAPost(t *testing.T) {
	t.Run("Successfully view a post", func(t *testing.T) {
		adapter := LocalBlogAdapter{}
		BlogAcceptanceCriteria(t, adapter)
	})
}
