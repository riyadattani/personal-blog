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

func NewServer(tempFolderPath string, repo Repository) (*mux.Router, error) {
	blogTemplate, err := template.ParseGlob(tempFolderPath)
	if err != nil {
		return nil, fmt.Errorf(
			"could not load template from %q, %v",
			tempFolderPath,
			err,
		)
	}

	server := server{
		blogTemplate: blogTemplate,
		repository:   repo,
	}

	router := mux.NewRouter()
	//TODO: At the moment this does not work
	//router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	router.HandleFunc("/", server.viewAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/about", server.viewAbout).Methods(http.MethodGet)
	router.HandleFunc("/blog/{title}", server.viewPost).Methods(http.MethodGet)

	return router, nil
}

func (s *server) viewAllPosts(writer http.ResponseWriter, request *http.Request) {
	s.blogTemplate.ExecuteTemplate(writer, "home.gohtml", s.repository.GetPosts())
}

func (s *server) viewAbout(writer http.ResponseWriter, request *http.Request) {
	s.blogTemplate.ExecuteTemplate(writer, "blog.gohtml", s.repository.GetPost("about.md"))
}

func (s *server) viewPost(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	title := vars["title"]
	s.blogTemplate.ExecuteTemplate(writer, "blog.gohtml", s.repository.GetPost(title))
}