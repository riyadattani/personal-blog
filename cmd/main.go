package main

import (
	"log"
	"personal-blog/pkg"
	"personal-blog/pkg/http_api"
)

func main() {
	app, err := pkg.NewApplication(pkg.NewConfig())
	if err != nil {
		log.Fatalf("Oops there is an error: %v", err)
	}

	server := http_api.NewServer(app.Config.ServerConfig, &app.Handler)

	log.Printf("listening on port %s\n", app.Config.Port)
	log.Fatal(server.ListenAndServe())
}
