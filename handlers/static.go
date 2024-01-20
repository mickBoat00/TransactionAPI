package handlers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	// Create a route pattern to match all files
	pattern := path + "/*"

	// Create the file server handler
	fs := http.StripPrefix(path, http.FileServer(root))

	// Register the file server handler with chi router
	r.Get(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
