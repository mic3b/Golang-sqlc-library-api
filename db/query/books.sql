-- name: CreateBook :one
INSERT INTO book (title, author, author_id) VALUES ($1, $2, $3) RETURNING *;

-- name: GetAllBooks :many
SELECT * FROM book;

-- name: GetBookByID :one
SELECT * FROM book WHERE id = $1;

-- name: UpdateBook :one
UPDATE book SET title = $1, author = $2, author_id = $3 WHERE id = $4 RETURNING *;

-- name: DeleteBook :one
DELETE FROM book WHERE id = $1 RETURNING *;