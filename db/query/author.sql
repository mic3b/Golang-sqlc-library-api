-- name: CreateAuthor :one
INSERT INTO author (first_name, last_name) VALUES ($1, $2) RETURNING *;

-- name: GetAllAuthors :many
SELECT * FROM author;