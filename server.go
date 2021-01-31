package blog

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Repository interface {
	GetPosts() []Post
	GetPost(title string) Post
}

type Server struct {
	blogTemplate *template.Template
	repository   Repository
	router       *mux.Router
}

func NewServer(tempFolderPath string, repo Repository) (*Server, error) {
	blogTemplate, err := template.ParseGlob(tempFolderPath)
	if err != nil {
		return nil, fmt.Errorf(
			"could not load template from %q, %v",
			tempFolderPath,
			err,
		)
	}
	router := mux.NewRouter()
	//TODO: At the moment this does not work
	//router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := blogTemplate.ExecuteTemplate(writer, "home.gohtml", repo.GetPosts())
		if err != nil {
			log.Fatal(fmt.Sprint("Could not execute blogTemplate", err))
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		err := blogTemplate.ExecuteTemplate(writer, "blog.gohtml", repo.GetPost("about.md"))
		if err != nil {
			log.Fatal(fmt.Sprint("Could not execute blogTemplate", err))
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/blog/{title}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["title"]
		err := blogTemplate.ExecuteTemplate(writer, "blog.gohtml", repo.GetPost(title))
		if err != nil {
			log.Fatal(fmt.Sprint("Could not execute blogTemplate", err))
		}
	}).Methods(http.MethodGet)

	return &Server{
		blogTemplate: blogTemplate,
		repository:   repo,
		router:       router,
	}, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
