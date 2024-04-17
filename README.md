# Helpers for the Gowebly CLI (and not so)

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

A most useful helpers for build the best **Go** web applications with [Gowebly CLI][gowebly_url].

> 💡 Note: You can use these helpers in other projects as well.

## 📖 List of helpers

### `gowebly.Getenv`

Helper to get the given environment variable. If key is not found, sets a fallback value.

```go
import (
    gowebly "github.com/gowebly/helpers"
)

// Get a value of the environment variable 'BACKEND_PORT'
// or sets it to a fallback value '5000'.
gowebly.Getenv("BACKEND_PORT", "5000")
```

> 💡 Note: This is a more advanced version of the built-in [os.Getenv][go_os_getenv_url] function.

### `gowebly.ParseTemplates`

Helper to parse list of the given templates to the HTTP handler.

```go
import (
    "log/slog"

    gowebly "github.com/gowebly/helpers"
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

### `gowebly.ParseTemplatesWithCustomMainLayout`

Helper to parse a list of the given templates with a custom main layout to the HTTP handler. Useful to use at times when you want to override file name of the default `templates/main.html` layout template.

```go
import (
    "log/slog"

    gowebly "github.com/gowebly/helpers"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Define path to the main layout template.
    customMainLayout := filepath.Join("templates", "my-custom-main.html")

    // Define paths to the user templates.
    indexPage := filepath.Join("templates", "pages", "index.html")
    indexLoginForm := filepath.Join("templates", "components", "index-login-form.html")

    // Parse user templates or return error.
    tmpl, err := gowebly.ParseTemplatesWithCustomMainLayout(customMainLayout, indexPage, indexLoginForm)
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

### `gowebly.StaticFileServerHandler`

Helpers to create a custom handler for serve embed `./static` folder.

```go
import (
    "embed"
    "net/http"

    gowebly "github.com/gowebly/helpers"
)

//go:embed static/*
var static embed.FS

// Create the gowebly helper for serve embed static folder.
staticFileServer := gowebly.StaticFileServerHandler(http.FS(static))

// Handle static files (with a custom handler).
http.Handle("/static/", staticFileServer)
```

## ⚠️ License

[`The Gowebly Helpers`][repo_url] is free and open-source software licensed under the [Apache 2.0 License][repo_license_url], created and supported by [Vic Shóstak][author_url] with 🩵 for people and robots.

<!-- Go links -->

[go_report_url]: https://goreportcard.com/report/github.com/gowebly/helpers
[go_dev_url]: https://pkg.go.dev/github.com/gowebly/helpers
[go_version_img]: https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go
[go_code_coverage_url]: https://codecov.io/gh/gowebly/helpers
[go_code_coverage_img]: https://img.shields.io/codecov/c/gh/gowebly/helpers.svg?logo=codecov&style=for-the-badge
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_os_getenv_url]: https://pkg.go.dev/os#Getenv

<!-- Repository links -->

[repo_url]: https://github.com/gowebly/helpers
[repo_license_url]: https://github.com/gowebly/helpers/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- README links -->

[gowebly_url]: https://github.com/gowebly/gowebly
