package create_url_test

import (
	"context"
	"testing"

	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/application/usecase/create_url"
	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/infra/persistence"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	t.Run("Should create a new URL", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		repository := persistence.NewMockUrlRepository(t)

		repository.
			EXPECT().
			Create(ctx, mock.Anything).
			Return(nil).
			Once()

		sut := create_url.New(repository)

		originalURL := "https://www.google.com"

		// Act
		url, err := sut.Execute(ctx, originalURL)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, url)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error when trying to create a new URL", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		repository := persistence.NewMockUrlRepository(t)

		repository.
			EXPECT().
			Create(ctx, mock.Anything).
			Return(assert.AnError).
			Once()

		sut := create_url.New(repository)

		originalURL := "https://www.google.com"

		// Act
		url, err := sut.Execute(ctx, originalURL)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, url)
		repository.AssertExpectations(t)
	})
}
