package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"personal-blog/pkg"
	"personal-blog/pkg/http_api"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	log.Println("listening on", port)

	if err := http.ListenAndServe(":"+port, newServer()); err != nil {
		log.Fatal("Cannot listen and serve", err)
	}
}

func newServer() *mux.Router {
	repository, err := pkg.NewInMemoryRepository(os.DirFS("posts"), os.DirFS("events"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create a repository: %s", err))
	}

	t, err := newTemplate("../html/*")
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create templates: %s", err))
	}

	server := http_api.NewHandler(t, repository)

	return http_api.NewRouter(server, "../css")
}

func newTemplate(tempFolderPath string) (*template.Template, error) {
	temp, err := template.ParseGlob(tempFolderPath)
	if err != nil {
		return nil, fmt.Errorf(
			"could not load template from %q, %v",
			tempFolderPath,
			err,
		)
	}
	return temp, nil
}
