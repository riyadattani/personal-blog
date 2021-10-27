package http_api

import (
	"net/http"
	"personal-blog/pkg/http_api/blog_handler"
)

func NewServer(config ServerConfig, handler *blog_handler.BlogHandler) (server *http.Server) {
	router := newRouter(handler, config.CSSDir)

	server = &http.Server{
		Addr:         config.TCPAddress(),
		Handler:      router,
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}

	return
}

