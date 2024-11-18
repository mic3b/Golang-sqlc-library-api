// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package query

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error)
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	DeleteBook(ctx context.Context, id int32) (Book, error)
	GetAllAuthors(ctx context.Context) ([]Author, error)
	GetAllBooks(ctx context.Context) ([]Book, error)
	GetBookByID(ctx context.Context, id int32) (Book, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error)
}
  
var _ Querier = (*Queries)(nil)
