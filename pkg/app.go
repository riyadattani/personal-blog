package pkg

import (
	"fmt"
	"html/template"
	"os"
	"personal-blog/pkg/http_api"
	"personal-blog/pkg/http_api/blog_handler"
	in_mem "personal-blog/pkg/in-mem"
	"time"
)

type App struct {
	Config  http_api.ServerConfig
	Handler blog_handler.BlogHandler
}

func NewApplication(config http_api.ServerConfig) (*App, error) {
	eventStore, err := in_mem.NewEventStore(os.DirFS(config.EventsDir))
	if err != nil {
		return nil, fmt.Errorf("failed to create the event store: %s", err)
	}

	postStore, err := in_mem.NewPostStore(os.DirFS(config.PostsDir))
	if err != nil {
		return nil, fmt.Errorf("failed to create the post store: %s", err)
	}

	templ, err := newTemplate(config.HTMLDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create templates: %s", err)
	}

	handler := blog_handler.NewHandler(templ, eventStore, postStore)

	return &App{
		Config:  config,
		Handler: *handler,
	}, nil
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

func NewConfig() http_api.ServerConfig {
	return http_api.ServerConfig{
		Port:             lookupEnvOr("PORT", defaultPort),
		HTTPReadTimeout:  defaultHTTPReadTimeout,
		HTTPWriteTimeout: defaulHTTPtWriteTimeout,
		CSSDir:           defaultCSSDir,
		HTMLDir:          defaultHTMLDir,
		PostsDir:         "posts",
		EventsDir:        "events",
	}
}

const (
	defaultCSSDir           = "../css"
	defaultHTMLDir          = "../html/*"
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
