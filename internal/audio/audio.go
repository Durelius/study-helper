// Package audio serves the spoken lecture notes.
//
// The files are not embedded in the binary: two and a half hours of speech is over a
// hundred megabytes, which would make every deploy ship the whole audiobook again to
// change a line of Go. They live in a directory instead, listed by the manifest that
// the render writes alongside them.
package audio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Episode is one lecture, spoken.
type Episode struct {
	// Topic matches the lecture id, so an episode can link to its notes.
	Topic   string  `json:"topic"`
	Index   int     `json:"index"`
	Title   string  `json:"title"`
	File    string  `json:"file"`
	Seconds float64 `json:"seconds"`
	Bytes   int64   `json:"bytes"`
	// URL is where the browser fetches it.
	URL string `json:"url"`
}

// Library is the rendered audiobook, or empty when none has been rendered.
type Library struct {
	dir      string
	Episodes []Episode
}

// Load reads the manifest from dir. A missing directory is not an error: the app runs
// perfectly well with no audio, and the tab simply says so.
func Load(dir string) (*Library, error) {
	lib := &Library{dir: dir, Episodes: []Episode{}}
	if dir == "" {
		return lib, nil
	}
	raw, err := os.ReadFile(filepath.Join(dir, "manifest.json"))
	if os.IsNotExist(err) {
		return lib, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(raw, &lib.Episodes); err != nil {
		return nil, fmt.Errorf("reading the audio manifest: %w", err)
	}
	for i := range lib.Episodes {
		e := &lib.Episodes[i]
		// A manifest entry with no file behind it would fail only when someone
		// pressed play, so check now.
		if _, err := os.Stat(filepath.Join(dir, e.File)); err != nil {
			return nil, fmt.Errorf("manifest lists %q but the file is missing", e.File)
		}
		// Episode filenames carry the lecture title, so they contain spaces and
		// ampersands. Browsers paper over a raw space; a correct URL does not need
		// them to.
		e.URL = "/media/" + url.PathEscape(e.File)
	}
	sort.Slice(lib.Episodes, func(i, j int) bool { return lib.Episodes[i].Index < lib.Episodes[j].Index })
	return lib, nil
}

// Minutes is the whole audiobook's length, for the tab to report.
func (l *Library) Minutes() float64 {
	total := 0.0
	for _, e := range l.Episodes {
		total += e.Seconds
	}
	return total / 60
}

// Handler serves the audio files themselves, under /media/.
//
// Deliberately not /audio/: that is the app's own page, and a Go mux pattern of
// "/audio/" swallows "/audio" with a redirect, leaving the SPA route unreachable.
//
// http.ServeFile handles range requests, which is what lets a player seek without
// downloading the whole episode first, and what makes Safari willing to play at all.
func (l *Library) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.URL.Path is already decoded, so the name here matches the manifest.
		name := strings.TrimPrefix(r.URL.Path, "/media/")
		if name == "" || strings.Contains(name, "/") || strings.Contains(name, "..") {
			http.NotFound(w, r)
			return
		}
		known := false
		for _, e := range l.Episodes {
			if e.File == name {
				known = true
				break
			}
		}
		if !known {
			http.NotFound(w, r)
			return
		}
		// The audio never changes without its name changing, and a phone on mobile
		// data should not refetch it.
		w.Header().Set("Cache-Control", "public, max-age=604800")
		http.ServeFile(w, r, filepath.Join(l.dir, name))
	})
}
