package blog_handler

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	in_mem "personal-blog/pkg/in-mem"
)

type BlogHandler struct {
	template   *template.Template
	postStore  in_mem.PostService
	eventStore in_mem.EventService
}

func NewHandler(
	template *template.Template,
	eventStore in_mem.EventService,
	postStore in_mem.PostService,
) *BlogHandler {
	return &BlogHandler{
		template:   template,
		postStore: postStore,
		eventStore: eventStore,
	}
}

func (s *BlogHandler) ViewAllPosts(w http.ResponseWriter, _ *http.Request) {
	err := s.template.ExecuteTemplate(w, "home.gohtml", s.postStore.GetPosts())
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
	post, err := s.postStore.GetPost(urlTitle)
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
	err := s.template.ExecuteTemplate(w, "events.gohtml", s.eventStore.GetEvents())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
