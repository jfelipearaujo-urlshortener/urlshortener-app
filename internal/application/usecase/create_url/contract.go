package create_url

import (
	"context"

	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/domain/entities"
)

type UseCase interface {
	Execute(ctx context.Context, originalURL string) (*entities.Url, error)
}
