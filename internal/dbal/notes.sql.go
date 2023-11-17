// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: notes.sql

package dbal

import (
	"context"
)

const createNote = `-- name: CreateNote :one
INSERT INTO notes (
  title, body
) VALUES (
  $1, $2
)
RETURNING id, title, body, version, created_at, updated_at
`

type CreateNoteParams struct {
	Title string `db:"title" json:"title"`
	Body  string `db:"body" json:"body"`
}

func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (Note, error) {
	row := q.db.QueryRow(ctx, createNote, arg.Title, arg.Body)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteNote = `-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1
`

func (q *Queries) DeleteNote(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteNote, id)
	return err
}

const getNote = `-- name: GetNote :one
SELECT id, title, body, version, created_at, updated_at FROM notes
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetNote(ctx context.Context, id int64) (Note, error) {
	row := q.db.QueryRow(ctx, getNote, id)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listNotes = `-- name: ListNotes :many
SELECT id, title, body, version, created_at, updated_at FROM notes
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListNotesParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) ListNotes(ctx context.Context, arg ListNotesParams) ([]Note, error) {
	rows, err := q.db.Query(ctx, listNotes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Note{}
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.Version,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateNote = `-- name: UpdateNote :exec
UPDATE notes
  SET title = $2, body = $3
WHERE id = $1
`

type UpdateNoteParams struct {
	ID    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Body  string `db:"body" json:"body"`
}

func (q *Queries) UpdateNote(ctx context.Context, arg UpdateNoteParams) error {
	_, err := q.db.Exec(ctx, updateNote, arg.ID, arg.Title, arg.Body)
	return err
}
