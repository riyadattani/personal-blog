package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"personal-blog/pkg"
	"personal-blog/pkg/http_api"
	"time"
)

func main() {
	if err := newServer().ListenAndServe(); err != nil {
		os.Exit(1)
	}
}

func newServer() *http.Server {
	repository, err := pkg.NewInMemoryRepository(os.DirFS("posts"), os.DirFS("events"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create a repository: %s", err))
	}

	t, err := newTemplate("../html/*")
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create templates: %s", err))
	}

	handler := http_api.NewHandler(t, repository)
	server := http_api.NewServer(newConfig(), handler, "../css")
	return server
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

func newConfig() http_api.ServerConfig {
	return http_api.ServerConfig{
		Port:             lookupEnvOr("PORT", defaultPort),
		HTTPReadTimeout:  defaultHTTPReadTimeout,
		HTTPWriteTimeout: defaulHTTPtWriteTimeout,
	}
}

const (
	defaultHTTPReadTimeout  = 2 * time.Second
	defaulHTTPtWriteTimeout = 2 * time.Second
	defaultPort             = "3000"
)

func lookupEnvOr(key string, defaultValue string) string {
	port, ok := os.LookupEnv(key)
	if !ok {
		port = defaultValue
	}
	return port
}
