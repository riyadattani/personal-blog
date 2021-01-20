package main

import (
	blog "personal-blog"
	"log"
	"net/http"
)

const addr = ":8080"

func main() {
	server, err := blog.NewServer("../../html/*", blog.NewInMemoryRepository())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on", addr)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatal("cannot listen and serve", err)
	}
}
