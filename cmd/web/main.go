package main

import (
	"log"
	"net/http"
	"os"
	"personal-blog/pkg"
	server2 "personal-blog/pkg/server"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	server, err := server2.NewServer(
		"../../html/*",
		"../../css",
		pkg.NewInMemoryRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on", port)
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatal("cannot listen and serve", err)
	}
}
