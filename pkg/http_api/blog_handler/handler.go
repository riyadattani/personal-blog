package blog_handler

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"personal-blog/pkg/ports"
)

type BlogHandler struct {
	template     *template.Template
	postService  ports.PostService
	eventService ports.EventService
}

func NewHandler(
	template *template.Template,
	eventService ports.EventService,
	postService ports.PostService,
) *BlogHandler {
	return &BlogHandler{
		template:     template,
		postService:  postService,
		eventService: eventService,
	}
}

func (s *BlogHandler) ViewAllPosts(w http.ResponseWriter, _ *http.Request) {
	err := s.template.ExecuteTemplate(w, "home.gohtml", s.postService.GetPosts())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *BlogHandler) ViewAbout(w http.ResponseWriter, _ *http.Request) {
	err := s.template.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *BlogHandler) ViewPost(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	urlTitle := vars["URLTitle"]
	post, err := s.postService.GetPost(urlTitle)
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

func (s *BlogHandler) ViewEvents(w http.ResponseWriter, e *http.Request) {
	err := s.template.ExecuteTemplate(w, "events.gohtml", s.eventService.GetEvents())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
