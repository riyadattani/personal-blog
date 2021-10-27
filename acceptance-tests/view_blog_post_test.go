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
type LocalBlogDriver struct {
}

func (l LocalBlogDriver) Publish(post blog.Post) error {

	//TODO: testing lib has a thing where it create a temp file and then deletes it when it is done - this is not working though - deleting before can read
	//postsDir, err := ioutil.TempDir(".", "posts-dir-test")
	//if err != nil {
	//	return err
	//}
	//defer os.RemoveAll(postsDir)

	//eventsDir, err := ioutil.TempDir(".", "events-dir-test")
	//if err != nil {
	//	return err
	//}
	//defer os.RemoveAll(eventsDir)

	//content := []byte(fmt.Sprintf("%s\n2013-Mar-03\npicture.jpg\ncat,dog\n-----\nThis is the first sentence of the content.\nThis is the second sentence.\n\nThis is the second paragraph.", post.Title))
	//tempFile, err := ioutil.TempFile(postsDirTest, "example_post")
	//if err != nil {
	//	return err
	//}
	//
	//_, err = tempFile.Write(content)
	//
	//if err != nil {
	//	return err
	//}

	file, err := os.Create("posts-test/example.md")
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

func (l LocalBlogDriver) ReadPost(urlTitle string) (bool, error) {
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
		driver := LocalBlogDriver{}
		BlogAcceptanceCriteria(t, driver)
	})
}
