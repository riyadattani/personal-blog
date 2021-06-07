package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"personal-blog/pkg"
	server "personal-blog/pkg/server"
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

	s, err := server.NewServer(
		"../html/*",
		"../css",
		repository,
	)
	if err != nil {
		log.Fatal(err)
	}

	return s
}

