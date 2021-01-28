package blog

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Repository interface {
	GetBlogs() []Blog
	GetBlog(title string) Blog
}

type Server struct {
	blogTemplate *template.Template
	repo         Repository
	router       *mux.Router
}

func NewServer(tempFolderPath string, repo Repository) (*Server, error) {
	blogTemplate, err := template.ParseGlob(tempFolderPath)
	if err != nil {
		return nil, fmt.Errorf(
			"could not load todo template from %q, %v",
			tempFolderPath,
			err,
		)
	}
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		blogTemplate.ExecuteTemplate(writer, "home.gohtml", repo.GetBlogs())
	}).Methods(http.MethodGet)

	router.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		blogTemplate.ExecuteTemplate(writer, "blog.gohtml", repo.GetBlog("about.md"))
	}).Methods(http.MethodGet)

	router.HandleFunc("/blog/{title}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["title"]
		blogTemplate.ExecuteTemplate(writer, "blog.gohtml", repo.GetBlog(title))
	}).Methods(http.MethodGet)

	return &Server{
		blogTemplate: blogTemplate,
		repo:         repo,
		router:       router,
	}, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
