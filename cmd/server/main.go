// Command server runs the study app for one course: its material, the quiz API and the
// leaderboard, all from one binary on one port.
//
// One process serves one course. Several courses means running this binary several
// times with a different -course, -db and port — which keeps a broken question file in
// one course from taking the others down with it, and keeps the code identical for all.
package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/wilhelmdurelius/chulastudy/internal/api"
	"github.com/wilhelmdurelius/chulastudy/internal/audio"
	"github.com/wilhelmdurelius/chulastudy/internal/content"
	"github.com/wilhelmdurelius/chulastudy/internal/course"
	"github.com/wilhelmdurelius/chulastudy/internal/store"
	"github.com/wilhelmdurelius/chulastudy/internal/web"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	courseDir := flag.String("course", "courses/valuechain", "directory holding this course's course.json, content and audio")
	dbPath := flag.String("db", "", "path to the SQLite file (default <course>/study.db)")
	audioDir := flag.String("audio", "", "directory holding the rendered audiobook (default <course>/audio)")
	examAt := flag.String("exam", "", "override the course's exam time, for the countdown")
	apiOnly := flag.Bool("dev", false, "serve only the API and let `vite dev` serve the UI")
	behindProxy := flag.Bool("behind-proxy", false, "a reverse proxy terminates TLS in front")
	flag.Parse()

	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))

	c, err := course.Load(*courseDir)
	if err != nil {
		log.Error("cannot load the course", "dir", *courseDir, "err", err)
		os.Exit(1)
	}
	if *examAt != "" {
		if c.ExamAt, err = time.Parse(time.RFC3339, *examAt); err != nil {
			log.Error("cannot read -exam", "err", err)
			os.Exit(1)
		}
	}

	// Loading the material up front turns an authoring mistake — an answer index out
	// of range, a topic this course does not teach — into a startup failure rather
	// than a wrong mark halfway through someone's revision.
	set, err := content.Load(c)
	if err != nil {
		log.Error("cannot load the study material", "course", c.ID, "err", err)
		os.Exit(1)
	}
	questions := 0
	for _, t := range set.Topics {
		questions += t.Questions
	}

	if *audioDir == "" {
		*audioDir = c.AudioDir()
	}
	lib, err := audio.Load(*audioDir)
	if err != nil {
		log.Error("cannot load the audiobook", "err", err)
		os.Exit(1)
	}

	if *dbPath == "" {
		*dbPath = filepath.Join(c.Dir, "study.db")
	}
	db, err := store.Open(*dbPath)
	if err != nil {
		log.Error("cannot open the database", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	log.Info("course loaded",
		"course", c.ID, "title", c.Title,
		"topics", len(set.Topics), "questions", questions,
		"lectures", len(set.Lectures), "cases", len(set.Cases),
		"glossary", len(set.Glossary), "episodes", len(lib.Episodes), "db", *dbPath)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	api.New(set, db, log, c.ExamAt, lib).Routes(mux)
	mux.Handle("GET /media/", lib.Handler())
	mux.Handle("GET /materials/", materials(c.MaterialsDir()))
	if !*apiOnly {
		mux.Handle("/", web.Handler())
	}

	srv := &http.Server{
		Addr:              *addr,
		Handler:           api.Secure(mux),
		ReadHeaderTimeout: 10 * time.Second,
	}
	_ = *behindProxy

	go func() {
		<-ctx.Done()
		shutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdown)
	}()

	log.Info("listening", "addr", *addr, "exam", c.ExamAt.Format(time.RFC1123))
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

// materials serves the rendered source slides a citation opens. A course without them
// answers 404 for the manifest, and the app then shows citations as plain text.
// Directory listings are refused so the route serves only the files the manifest names.
func materials(dir string) http.Handler {
	files := http.StripPrefix("/materials/", http.FileServer(http.Dir(dir)))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		// Re-rendering keeps the names, so cache for a while but not forever.
		w.Header().Set("Cache-Control", "public, max-age=3600")
		files.ServeHTTP(w, r)
	})
}
