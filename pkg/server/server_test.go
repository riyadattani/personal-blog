package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"personal-blog/pkg/blog"
	"strings"
	"testing"
)

type StubRepo struct {
	posts []blog.Post
}

func (s *StubRepo) GetPosts() []blog.Post {
	return s.posts
}

func (s *StubRepo) GetPost(title string) (blog.Post, error) {
	for _, post := range s.posts {
		if post.Title == title {
			return post, nil
		}
	}

	return blog.Post{}, fmt.Errorf("could not find %q in %+v", title, s.posts)
}

func TestServer(t *testing.T) {
	post := blog.Post{
		Title:   "this is a title",
		Content: "HTML template which is basically a string",
	}

	post2 := blog.Post{
		Title:   "this is another title",
		Content: "HTML template which is basically a string",
	}

	repo := StubRepo{
		[]blog.Post{post, post2},
	}

	template, err := newBlogTemplate("../../html/*")
	if err != nil {
		t.Fatal("could not load blog template", err)
	}

	blogServer := &BlogServer{
		blogTemplate: template,
		repository:   &repo,
	}

	t.Run("returns status code 200 on home page when getting posts", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		blogServer.viewAllPosts(response, request)

		gotStatusCode := response.Code
		wantStatusCode := http.StatusOK

		body := response.Body.String()

		if gotStatusCode != wantStatusCode {
			t.Fatalf("got %d, want %d", gotStatusCode, wantStatusCode)
		}

		if !strings.Contains(body, post.Title) {
			t.Error("Response body does not contain the first post")
		}

		if !strings.Contains(body, post2.Title) {
			t.Error("Response body does not contain the second post")
		}
	})

	t.Run("returns a status OK on a single post and has the content", func(t *testing.T) {
		server2, _ := NewServer("../../html/*", "../../css/*", &repo)
		newServer := httptest.NewServer(server2)
		defer newServer.Close()

		url := newServer.URL + "/blog/this is a title"

		res, err := http.Get(url)
		if err != nil {
			t.Fatal(err)
		}

		gotStatusCode := res.StatusCode
		wantStatusCode := http.StatusOK

		if gotStatusCode != wantStatusCode {
			t.Fatalf("got %d, want %d", gotStatusCode, wantStatusCode)
		}

		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(body), string(post.Content)) {
			t.Error("Response body does not contain the first post content")
		}
	})

	t.Run("returns a status OK on the about page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/about", nil)
		response := httptest.NewRecorder()

		blogServer.viewAbout(response, request)

		gotStatusCode := response.Code
		wantStatusCode := http.StatusOK

		if gotStatusCode != wantStatusCode {
			t.Fatalf("got %d, want %d", gotStatusCode, wantStatusCode)
		}
	})
}
