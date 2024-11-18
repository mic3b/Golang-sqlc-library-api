package services

import (
	"context"

	"example.com/m/initializer"
	"example.com/m/internal/model"
	"example.com/m/query"
	"example.com/m/utils"
)

type AuthorService struct {
	repo *query.Queries
}

func NewAuthorService() *AuthorService {
	return &AuthorService{
		repo: initializer.InitQuery(),
	}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, params query.CreateAuthorParams) (author model.Author, err error) {
	authorDB, nil := s.repo.CreateAuthor(ctx, params)
	if err != nil {
		return model.Author{}, err
	}

	utils.Convert(authorDB, &author)

	return author, nil
}

func (s *AuthorService) GetAllAuthors(ctx context.Context) (authors []model.Author, err error) {
	authorsDB, err := s.repo.GetAllAuthors(ctx)
	if err != nil {
		return nil, err
	}

	utils.Convert(authorsDB, &authors)

	return authors, nil
}
