package http_api
//
//import (
//	"fmt"
//	is2 "github.com/matryer/is"
//	"net/http"
//	"net/http/httptest"
//	"personal-blog/pkg/blog"
//	"personal-blog/pkg/event"
//	"strings"
//	"testing"
//)
//
//type StubRepo struct {
//	posts  []blog.Post
//	events []event.Event
//}
//
//func (s *StubRepo) GetPosts() []blog.Post {
//	return s.posts
//}
//
//func (s *StubRepo) GetPost(title string) (blog.Post, error) {
//	for _, post := range s.posts {
//		if post.Title == title {
//			return post, nil
//		}
//	}
//
//	return blog.Post{}, fmt.Errorf("could not find %q in %+v", title, s.posts)
//}
//
//func (s *StubRepo) GetEvents() []event.Event {
//	return s.events
//}
//
//func TestServer(t *testing.T) {
//	is := is2.New(t)
//
//	post := blog.Post{
//		URLTitle: "this-is-a-title",
//		Title:   "this is a title",
//		Content: "HTML template which is basically a string",
//	}
//
//	post2 := blog.Post{
//		URLTitle: "this-is-another-title",
//		Title:   "this is another title",
//		Content: "HTML template which is basically a string",
//	}
//
//	event1 := event.Event{Title: "Event1"}
//	event2 := event.Event{Title: "Event2"}
//
//	repo := StubRepo{
//		[]blog.Post{post, post2},
//		[]event.Event{event1, event2},
//	}
//
//	template, err := newTemplate("../../html/*")
//	if err != nil {
//		t.Fatal("could not load blog template", err)
//	}
//
//	blogServer := &BlogHandler{
//		template:   template,
//		repository: &repo,
//	}
//
//	t.Run("Posts", func(t *testing.T) {
//		t.Run("returns status code 200 on home page when getting posts", func(t *testing.T) {
//			request, _ := http.NewRequest(http.MethodGet, "/", nil)
//			response := httptest.NewRecorder()
//
//			blogServer.viewAllPosts(response, request)
//
//			body := response.Body.String()
//
//			is.Equal(response.Code, http.StatusOK)
//
//			//TODO: this does not work in the test (works on live) because we are not using newRouter in this test. Source of header code: https://stackoverflow.com/questions/51456253/how-to-set-http-responsewriter-content-type-header-globally-for-all-api-endpoint
//			//if response.Header().Get("Cache-Control") != "public, max-age=86400" {
//			//	t.Error("Response header does not contain the cache control values")
//			//}
//
//			if !strings.Contains(body, post.Title) {
//				t.Error("Response body does not contain the first post")
//			}
//
//			if !strings.Contains(body, post2.Title) {
//				t.Error("Response body does not contain the second post")
//			}
//		})
//
//		//TODO: MAJOR TODO - why is this not working?
//		//t.Run("returns a status OK on a single post and has the content", func(t *testing.T) {
//		//	request, _ := http.NewRequest(http.MethodGet, "/blog/this-is-a-title", nil)
//		//	response := httptest.NewRecorder()
//		//
//		//	blogServer.viewPost(response, request)
//		//
//		//	is.Equal(response.Code, http.StatusOK)
//		//
//		//	body := response.Body.String()
//		//
//		//	if !strings.Contains(body, string(post.Content)) {
//		//		t.Error("Response body does not contain the first post content")
//		//	}
//		//
//		//	//if response.Header.Get("Cache-Control") != "public, max-age=86400" {
//		//	//	t.Error("Response header does not contain the cache control values")
//		//	//}
//		//})
//	})
//
//	t.Run("Events", func(t *testing.T) {
//		request, _ := http.NewRequest(http.MethodGet, "/events", nil)
//		response := httptest.NewRecorder()
//
//		blogServer.viewEvents(response, request)
//
//		body := response.Body.String()
//
//		is.Equal(response.Code, http.StatusOK)
//
//		if !strings.Contains(body, event1.Title) {
//			t.Error("Response body does not contain the first event")
//		}
//
//		if !strings.Contains(body, event2.Title) {
//			t.Error("Response body does not contain the second event")
//		}
//	})
//
//	t.Run("returns a status OK on the about page", func(t *testing.T) {
//		request, _ := http.NewRequest(http.MethodGet, "/about", nil)
//		response := httptest.NewRecorder()
//
//		blogServer.viewAbout(response, request)
//
//		is.Equal(response.Code, http.StatusOK)
//
//		//TODO: this does not work in the test (works on live) because we are not using newRouter in this test. Source of header code: https://stackoverflow.com/questions/51456253/how-to-set-http-responsewriter-content-type-header-globally-for-all-api-endpoint
//		//if response.Header().Get("Cache-Control") != "public, max-age=86400" {
//		//	t.Error("Response header does not contain the cache control values")
//		//}
//	})
//}
