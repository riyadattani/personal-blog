package blog

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Repository interface {
	GetPosts() []Post
	GetPost(title string) Post
}

type server struct {
	blogTemplate *template.Template
	repository   Repository
}

func NewServer(tempFolderPath string, cssFolderPath string, repo Repository) (*mux.Router, error) {
	blogTemplate, err := newBlogTemplate(tempFolderPath)
	if err != nil {
		return nil, err
	}

	server := server{
		blogTemplate: blogTemplate,
		repository:   repo,
	}

	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(cssFolderPath))))
	router.HandleFunc("/", server.viewAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/about", server.viewAbout).Methods(http.MethodGet)
	router.HandleFunc("/blog/{title}", server.viewPost).Methods(http.MethodGet)

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

func (s *server) viewAllPosts(w http.ResponseWriter, _ *http.Request) {
	s.blogTemplate.ExecuteTemplate(w, "home.gohtml", s.repository.GetPosts())
}

func (s *server) viewAbout(w http.ResponseWriter, _ *http.Request) {
	s.blogTemplate.ExecuteTemplate(w, "about.gohtml", nil)
}

func (s *server) viewPost(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	title := vars["title"]
	s.blogTemplate.ExecuteTemplate(w, "blog.gohtml", s.repository.GetPost(title))
}
