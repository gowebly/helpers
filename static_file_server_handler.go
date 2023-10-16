package gowebly

import (
	"log/slog"
	"net/http"
)

// StaticFileServerHandler handles a custom handler for serve embed static folder.
//
// Example:
//
//	import (
//		"embed"
//		"net/http"
//
//		gowebly "github.com/gowebly/helpers"
//	)
//
//	//go:embed static/*
//	var static embed.FS
//
//	// Handle static files (with a custom handler).
//	http.Handle("/static/", gowebly.StaticFileServerHandler(http.FS(static)))
func StaticFileServerHandler(fs http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the requested file exists.
		_, err := fs.Open(r.URL.Path)
		if err != nil {
			// If the file is not found, return HTTP 404 error.
			http.NotFound(w, r)
			slog.Error(err.Error(), "method", r.Method, "status", http.StatusNotFound, "path", r.URL.Path)
			return
		}

		// File is found, serve it using the standard http.FileServer.
		http.FileServer(fs).ServeHTTP(w, r)
	})
}
