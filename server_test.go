package blog

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
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

	t.Run("returns a list of blog posts", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.viewAllPosts(response, request)

		got := response.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

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
