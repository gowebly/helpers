# gowebly helpers

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][repo_url]
[![License][repo_license_img]][repo_license_url]

A most useful helpers for build the best **Go** web applications with 
[`gowebly`][gowebly_url] CLI.

## 📖 List of helpers

### `gowebly.ParseTemplates`

Helper to parse list of the given templates to the HTTP handler.

```go
import (
    "log/slog"
    
    "github.com/gowebly/helpers"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Define paths to the user templates.
    indexPage := filepath.Join("templates", "pages", "index.html")
    indexLoginForm := filepath.Join("templates", "components", "index-login-form.html")
    
    // Parse user templates, using gowebly helper, or return error.
    tmpl, err := gowebly.ParseTemplates(indexPage, indexLoginForm)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        slog.Error(err.Error(), "method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path)
        return
    }
    
    // Execute (render) all templates or return error.
    if err := tmpl.Execute(w, nil); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        slog.Error(err.Error(), "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
        return
    }
}
```

> 💡 Note: The main layout template (`templates/main.html`) is already included.

### `gowebly.StaticFileServerHandler`

Helpers to create a custom handler for serve embed `./static` folder.

```go
import (
	"embed"
	"net/http"

	"github.com/gowebly/helpers"
)

//go:embed static/*
var static embed.FS

// Create the gowebly helper for serve embed static folder.
staticFileServer := gowebly.StaticFileServerHandler(http.FS(static))

// Handle static files (with a custom handler).
http.Handle("/static/", staticFileServer)
```

## ⚠️ License

[`gowebly helpers`][repo_url] is free and open-source software licensed 
under the [Apache 2.0 License][repo_license_url], created and supported by 
[Vic Shóstak][author_url] with 🩵 for people and robots.

<!-- Go links -->

[go_report_url]: https://goreportcard.com/report/github.com/gowebly/helpers
[go_dev_url]: https://pkg.go.dev/github.com/gowebly/helpers
[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[go_code_coverage_img]: https://img.shields.io/badge/code_coverage-0%25-success?style=for-the-badge&logo=none
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none

<!-- Repository links -->

[repo_url]: https://github.com/gowebly/helpers
[repo_license_url]: https://github.com/gowebly/helpers/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- README links -->

[gowebly_url]: https://github.com/gowebly/gowebly
