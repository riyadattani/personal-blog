package http_api

import (
	"github.com/gorilla/mux"
	"net/http"
	"personal-blog/pkg/http_api/blog_handler"
)

func newRouter(handler *blog_handler.BlogHandler, cssFolderPath string) *mux.Router {
	//TODO: allow cloudflare to cache website
	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Cache-Control", "public, max-age=86400")
			next.ServeHTTP(w, r)
		})
	})
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(cssFolderPath))))
	router.HandleFunc("/", handler.ViewAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/about", handler.ViewAbout).Methods(http.MethodGet)
	router.HandleFunc("/blog/{URLTitle}", handler.ViewPost).Methods(http.MethodGet)
	router.HandleFunc("/events", handler.ViewEvents).Methods(http.MethodGet)
	//TODO: You can create a custom 404 page
	router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	return router
}
