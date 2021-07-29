package pkg

import (
	"fmt"
	"html/template"
	"os"
	"personal-blog/pkg/http_api"
	"time"
)

type App struct {
	Config  http_api.ServerConfig
	Handler http_api.BlogHandler
}

func NewApplication() (*App, error) {
	config := newConfig()
	repository, err := NewInMemoryRepository(os.DirFS("posts"), os.DirFS("events"))
	if err != nil {
		return nil, fmt.Errorf("failed to create a repository: %s", err)
	}

	templ, err := newTemplate(config.HTMLDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create templates: %s", err)
	}

	handler := http_api.NewHandler(templ, repository)

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

func newConfig() http_api.ServerConfig {
	return http_api.ServerConfig{
		Port:             lookupEnvOr("PORT", defaultPort),
		HTTPReadTimeout:  defaultHTTPReadTimeout,
		HTTPWriteTimeout: defaulHTTPtWriteTimeout,
		CSSDir:           defaultCSSDir,
		HTMLDir:          defaultHTMLDir,
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
