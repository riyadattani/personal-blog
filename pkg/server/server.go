package server

import (
	"fmt"
	"html/template"
	"net/http"
	"personal-blog/pkg/blog"

	"github.com/gorilla/mux"
)

type Repository interface {
	GetPosts() []blog.Post
	GetPost(title string) (blog.Post, error)
}

type BlogServer struct {
	blogTemplate *template.Template
	repository   Repository
}

func NewServer(tempFolderPath string, cssFolderPath string, repo Repository) (*mux.Router, error) {
	blogTemplate, err := newBlogTemplate(tempFolderPath)
	if err != nil {
		return nil, err
	}

	server := BlogServer{
		blogTemplate: blogTemplate,
		repository:   repo,
	}
	//TODO: allow cloudflare to cache website
	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(cssFolderPath))))
	router.HandleFunc("/", server.viewAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/about", server.viewAbout).Methods(http.MethodGet)
	router.HandleFunc("/blog/{title}", server.viewPost).Methods(http.MethodGet)
	//TODO: You can create a custom 404 page
	router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	return router, nil
}

func newBlogTemplate(tempFolderPath string) (*template.Template, error) {
	blogTemplate, err := template.ParseGlob(tempFolderPath)
	if err != nil {
		return nil, fmt.Errorf(
			"could not load template from %q, %v",
			tempFolderPath,
			err,
		)
	}
	return blogTemplate, nil
}

func (s *BlogServer) viewAllPosts(w http.ResponseWriter, _ *http.Request) {
	s.blogTemplate.ExecuteTemplate(w, "home.gohtml", s.repository.GetPosts())
}

func (s *BlogServer) viewAbout(w http.ResponseWriter, _ *http.Request) {
	s.blogTemplate.ExecuteTemplate(w, "about.gohtml", nil)
}

func (s *BlogServer) viewPost(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	title := vars["title"]
	post, err := s.repository.GetPost(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	s.blogTemplate.ExecuteTemplate(w, "blog.gohtml", post)
}
