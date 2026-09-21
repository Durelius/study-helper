// Package web serves the built React app from the binary.
package web

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

// dist is filled by `npm run build` in web/, which writes here directly. The
// directory carries a committed .gitkeep so this compiles on a fresh checkout,
// before the app has ever been built.
//
//go:embed all:dist
var dist embed.FS

// Handler serves the built app with an SPA fallback, so client-side routes survive a
// hard refresh.
func Handler() http.Handler {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
	files := http.FileServer(http.FS(sub))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}
		if _, err := fs.Stat(sub, path); err != nil {
			// Unknown path with no file extension: hand it to the SPA router.
			r = r.Clone(r.Context())
			r.URL.Path = "/"
			path = "index.html"
		}

		// Asset filenames carry a content hash, so they can be cached forever. The
		// page that points at them must not be, or a rebuild leaves the browser
		// running the previous app against the new API.
		if strings.HasPrefix(path, "assets/") {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		} else {
			w.Header().Set("Cache-Control", "no-cache")
		}
		files.ServeHTTP(w, r)
	})
}
