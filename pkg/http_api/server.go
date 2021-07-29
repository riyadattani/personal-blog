package http_api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewServer(config ServerConfig, handler *BlogHandler) (server *http.Server) {
	router := newRouter(handler, config.CSSDir)

	server = &http.Server{
		Addr:         config.TCPAddress(),
		Handler:      router,
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}

	return
}

func newRouter(handler *BlogHandler, cssFolderPath string) *mux.Router {
	//TODO: allow cloudflare to cache website
	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Cache-Control", "public, max-age=86400")
			next.ServeHTTP(w, r)
		})
	})
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(cssFolderPath))))
	router.HandleFunc("/", handler.viewAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/about", handler.viewAbout).Methods(http.MethodGet)
	router.HandleFunc("/blog/{URLTitle}", handler.viewPost).Methods(http.MethodGet)
	router.HandleFunc("/events", handler.viewEvents).Methods(http.MethodGet)
	//TODO: You can create a custom 404 page
	router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	return router
}
