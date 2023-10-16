package gowebly

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// ParseTemplates parses list of the given templates to the HTTP handler.
//
// Already included 'templates/main.html' layout template from your project
// path.
//
// Example:
//
//	import (
//		"log/slog"
//
//		gowebly "github.com/gowebly/helpers"
//	)
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//		// Define paths to the user templates.
//		indexPage := filepath.Join("templates", "pages", "index.html")
//		indexLoginForm := filepath.Join("templates", "components", "index-login-form.html")
//
//		// Parse user templates or return error.
//		tmpl, err := gowebly.ParseTemplates(indexPage, indexLoginForm)
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			slog.Error(err.Error(), "method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path)
//			return
//		}
//
//		// Execute (render) all templates or return error.
//		if err := tmpl.Execute(w, nil); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			slog.Error(err.Error(), "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
//			return
//		}
//	}
func ParseTemplates(names ...string) (*template.Template, error) {
	// Set global templates.
	global := []string{
		filepath.Join("templates", "main.html"),
	}

	// Check if all templates exist.
	for _, n := range names {
		if !isExistInFolder(n, false) {
			return nil, fmt.Errorf("os: template '%s' is not found", n)
		}
	}

	// Add all user templates after global.
	global = append(global, names...)

	return template.ParseFiles(global...)
}

// ParseTemplatesWithCustomMainLayout parses list of the given templates with a
// custom main layout to the HTTP handler.
//
// Example:
//
//	import (
//		"log/slog"
//
//		gowebly "github.com/gowebly/helpers"
//	)
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//		// Define path to the main layout template.
//		customMainLayout := filepath.Join("templates", "my-custom-main.html")
//
//		// Define paths to the user templates.
//		indexPage := filepath.Join("templates", "pages", "index.html")
//		indexLoginForm := filepath.Join("templates", "components", "index-login-form.html")
//
//		// Parse user templates or return error.
//		tmpl, err := gowebly.ParseTemplatesWithCustomMainLayout(customMainLayout, indexPage, indexLoginForm)
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			slog.Error(err.Error(), "method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path)
//			return
//		}
//
//		// Execute (render) all templates or return error.
//		if err := tmpl.Execute(w, nil); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			slog.Error(err.Error(), "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
//			return
//		}
//	}
func ParseTemplatesWithCustomMainLayout(main string, names ...string) (*template.Template, error) {
	// Set global templates.
	global := make([]string, 0, len(names)+1)
	global = append(global, main)

	for _, n := range names {
		// Check, if the given template is existing.
		if !isExistInFolder(n, false) {
			return nil, fmt.Errorf("os: template '%s' is not found", n)
		}

		// Add the given template after global.
		global = append(global, n)
	}

	return template.Must(template.ParseFiles(global...)), nil
}

// isExistInFolder searches for a file or folder by the given name in the
// current folder.
func isExistInFolder(name string, isFolder bool) bool {
	// Check if file or folder exists.
	_, err := os.Stat(filepath.Clean(name))
	if err != nil {
		return false
	}

	// Check if it is a directory.
	info, err := os.Lstat(filepath.Clean(name))
	if err != nil {
		return false
	}

	return info.IsDir() == isFolder
}
