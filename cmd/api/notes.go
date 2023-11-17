package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/mymorkkis/notes-app/internal/dbal"
)

func (app *application) NoteCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noteID, err := strconv.ParseInt(chi.URLParam(r, "noteID"), 10, 0)
		if err != nil {
			app.errorResponse(w, r, http.StatusBadRequest, "Note ID is not a valid integer")
			return
		}

		note, err := app.queries.GetNote(r.Context(), noteID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				app.errorResponse(w, r, http.StatusBadRequest, "Note could not be found")
			} else {
				app.internalServerError(w, r, err.Error())
			}
			return
		}

		ctx := context.WithValue(r.Context(), "note", note)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) getNote(w http.ResponseWriter, r *http.Request) {
	note, ok := r.Context().Value("note").(dbal.Note)
	if !ok {
		app.internalServerError(w, r, "Cannot fetch note from context")
		return
	}

	serveJSON(w, r, http.StatusOK, note, app.systemInfo, nil)
}

func (app *application) createNote(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	err := readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	note, err := app.queries.CreateNote(r.Context(), dbal.CreateNoteParams(input))
	if err != nil {
		app.internalServerError(w, r, err.Error())
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/notes/%d", note.ID))

	serveJSON(w, r, http.StatusCreated, note, app.systemInfo, headers)
}
