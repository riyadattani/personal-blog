package main

import (
	"log"
	"net/http"
	blog "personal-blog"
)

const addr = ":3000"

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
