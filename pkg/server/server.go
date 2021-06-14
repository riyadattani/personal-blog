package server

import (
	"fmt"
	"html/template"
	"net/http"
	"personal-blog/pkg/blog"
	"personal-blog/pkg/event"

	"github.com/gorilla/mux"
)

type Repository interface {
	GetPosts() []blog.Post
	GetPost(title string) (blog.Post, error)
	GetEvents() []event.Event
}

type BlogServer struct {
	template   *template.Template
	repository Repository
}

func NewServer(tempFolderPath string, cssFolderPath string, repo Repository) (*mux.Router, error) {
	templ, err := newBlogTemplate(tempFolderPath)
	if err != nil {
		return nil, err
	}

	server := BlogServer{
		template:   templ,
		repository: repo,
	}
	//TODO: allow cloudflare to cache website
	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Cache-Control", "public, max-age=86400")
			next.ServeHTTP(w, r)
		})
	})
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(cssFolderPath))))
	router.HandleFunc("/", server.viewAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/about", server.viewAbout).Methods(http.MethodGet)
	router.HandleFunc("/blog/{URLTitle}", server.viewPost).Methods(http.MethodGet)
	router.HandleFunc("/events", server.viewEvents).Methods(http.MethodGet)
	//TODO: You can create a custom 404 page
	router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	return router, nil
}

func newBlogTemplate(tempFolderPath string) (*template.Template, error) {
	template, err := template.ParseGlob(tempFolderPath)
	if err != nil {
		return nil, fmt.Errorf(
			"could not load template from %q, %v",
			tempFolderPath,
			err,
		)
	}
	return template, nil
}

func (s *BlogServer) viewAllPosts(w http.ResponseWriter, _ *http.Request) {
	s.template.ExecuteTemplate(w, "home.gohtml", s.repository.GetPosts())
}

func (s *BlogServer) viewAbout(w http.ResponseWriter, _ *http.Request) {
	s.template.ExecuteTemplate(w, "about.gohtml", nil)
}

func (s *BlogServer) viewPost(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	urlTitle := vars["URLTitle"]
	post, err := s.repository.GetPost(urlTitle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	s.template.ExecuteTemplate(w, "blog.gohtml", post)
}

func (s *BlogServer) viewEvents(w http.ResponseWriter, e *http.Request) {
	s.template.ExecuteTemplate(w, "events.gohtml", s.repository.GetEvents())
}
