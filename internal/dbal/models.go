// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package dbal

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Note struct {
	ID        int64              `db:"id" json:"id"`
	Title     string             `db:"title" json:"title"`
	Body      string             `db:"body" json:"body"`
	Version   int32              `db:"version" json:"version"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}
