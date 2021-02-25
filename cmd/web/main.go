package main

import (
	"log"
	"net/http"
	"personal-blog/pkg"
	server2 "personal-blog/pkg/server"
)

const addr = ":3000"

func main() {
	server, err := server2.NewServer(
		"../../html/*",
		"../../css",
		pkg.NewInMemoryRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on", addr)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatal("cannot listen and serve", err)
	}
}
