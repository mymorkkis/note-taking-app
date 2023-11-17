-- name: GetNote :one
SELECT * FROM notes
WHERE id = $1 LIMIT 1;

-- name: ListNotes :many
SELECT * FROM notes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateNote :one
INSERT INTO notes (
  title, body
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateNote :exec
UPDATE notes
  SET title = $2, body = $3
WHERE id = $1;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;
