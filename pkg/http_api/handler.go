package http_api

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"personal-blog/pkg/blog"
	"personal-blog/pkg/event"
)

type Repository interface {
	GetPosts() []blog.Post
	GetPost(title string) (blog.Post, error)
	GetEvents() []event.Event
}

type BlogHandler struct {
	template   *template.Template
	repository Repository
}

func NewHandler(template *template.Template, repo Repository) *BlogHandler {
	return &BlogHandler{
		template:   template,
		repository: repo,
	}
}

func (s *BlogHandler) viewAllPosts(w http.ResponseWriter, _ *http.Request) {
	err := s.template.ExecuteTemplate(w, "home.gohtml", s.repository.GetPosts())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *BlogHandler) viewAbout(w http.ResponseWriter, _ *http.Request) {
	err := s.template.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *BlogHandler) viewPost(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	urlTitle := vars["URLTitle"]
	post, err := s.repository.GetPost(urlTitle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = s.template.ExecuteTemplate(w, "blog.gohtml", post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *BlogHandler) viewEvents(w http.ResponseWriter, e *http.Request) {
	err := s.template.ExecuteTemplate(w, "events.gohtml", s.repository.GetEvents())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}