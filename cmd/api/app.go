package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mymorkkis/notes-app/internal/dbal"
)

type systemInfo struct {
	version     string
	environment string
}

type application struct {
	systemInfo systemInfo
	logger     *slog.Logger
	queries    *dbal.Queries
}

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(middleware.AllowContentType("application/json"))

	r.Route("/notes", func(r chi.Router) {
		r.Get("/", app.listNotes)
		r.Post("/", app.createNote)
		r.Route("/{noteID}", func(r chi.Router) {
			r.Use(app.NoteCtx)
			r.Get("/", app.getNote)
		})
	})

	return r
}

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, message string) {
	app.logger.Error(message)
	status := http.StatusInternalServerError
	app.errorResponse(w, r, status, http.StatusText(status))
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	data := map[string]any{"error": message}
	serveJSON(w, r, status, data, app.systemInfo, nil)
}
