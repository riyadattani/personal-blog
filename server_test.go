package blog

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubRepo struct {
	posts []Post
	post  Post
}

func (s *StubRepo) GetPosts() []Post {
	return s.posts
}

func (s *StubRepo) GetPost(title string) Post {
	return s.post
}

func TestServer(t *testing.T) {
	post := Post{
		Title:   "this is a title",
		Content: "HTML template which is basically a string",
	}

	post2 := Post{
		Title:   "this is another title",
		Content: "HTML template which is basically a string",
	}

	repo := StubRepo{
		[]Post{post, post2},
		post,
	}

	template, err := newBlogTemplate("./html/*")
	if err != nil {
		t.Fatal("could not load blog template", err)
	}

	server := &server{
		blogTemplate: template,
		repository:   &repo,
	}

	t.Run("returns status code 200 on home page when getting posts", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.viewAllPosts(response, request)

		gotStatusCode := response.Code
		wantStatusCode := http.StatusOK

		if gotStatusCode != wantStatusCode {
			t.Errorf("got %q, want %q", gotStatusCode, wantStatusCode)
		}
	})

	t.Run("returns a status OK on a single post", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/blog/this is another title", nil)
		response := httptest.NewRecorder()

		server.viewPost(response, request)

		gotStatusCode := response.Code
		wantStatusCode := http.StatusOK

		if gotStatusCode != wantStatusCode {
			t.Errorf("got %q, want %q", gotStatusCode, wantStatusCode)
		}
	})

	t.Run("returns a status OK on the about page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/about", nil)
		response := httptest.NewRecorder()

		server.viewAbout(response, request)

		gotStatusCode := response.Code
		wantStatusCode := http.StatusOK

		if gotStatusCode != wantStatusCode {
			t.Errorf("got %q, want %q", gotStatusCode, wantStatusCode)
		}
	})
}

