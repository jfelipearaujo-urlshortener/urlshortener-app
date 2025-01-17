package create_url

import (
	"context"

	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/domain/entities"
	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/infra/persistence"
)

type useCase struct {
	repository persistence.UrlRepository
}

func New(repository persistence.UrlRepository) *useCase {
	return &useCase{
		repository: repository,
	}
}

func (usc *useCase) Execute(ctx context.Context, originalURL string) (*entities.Url, error) {
	url := entities.NewURL(originalURL, "asda12831")

	if err := usc.repository.Create(ctx, url); err != nil {
		return nil, err
	}

	return url, nil
}
