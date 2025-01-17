package persistence

import (
	"context"

	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/domain/entities"
)

type UrlRepository interface {
	Create(ctx context.Context, url *entities.Url) error
}
