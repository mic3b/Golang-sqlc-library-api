// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: books.sql

package query

import (
	"context"
	"database/sql"
)

const createBook = `-- name: CreateBook :one
INSERT INTO book (title, author, author_id) VALUES ($1, $2, $3) RETURNING id, title, author, created_at, author_id
`

type CreateBookParams struct {
	Title    string
	Author   string
	AuthorID sql.NullInt32
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook, arg.Title, arg.Author, arg.AuthorID)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.CreatedAt,
		&i.AuthorID,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :one
DELETE FROM book WHERE id = $1 RETURNING id, title, author, created_at, author_id
`

func (q *Queries) DeleteBook(ctx context.Context, id int32) (Book, error) {
	row := q.db.QueryRowContext(ctx, deleteBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.CreatedAt,
		&i.AuthorID,
	)
	return i, err
}

const getAllBooks = `-- name: GetAllBooks :many
SELECT id, title, author, created_at, author_id FROM book
`

func (q *Queries) GetAllBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, getAllBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.CreatedAt,
			&i.AuthorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookByID = `-- name: GetBookByID :one
SELECT id, title, author, created_at, author_id FROM book WHERE id = $1
`

func (q *Queries) GetBookByID(ctx context.Context, id int32) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBookByID, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.CreatedAt,
		&i.AuthorID,
	)
	return i, err
}

const updateBook = `-- name: UpdateBook :one
UPDATE book SET title = $1, author = $2, author_id = $3 WHERE id = $4 RETURNING id, title, author, created_at, author_id
`

type UpdateBookParams struct {
	Title    string
	Author   string
	AuthorID sql.NullInt32
	ID       int32
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBook,
		arg.Title,
		arg.Author,
		arg.AuthorID,
		arg.ID,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.CreatedAt,
		&i.AuthorID,
	)
	return i, err
}