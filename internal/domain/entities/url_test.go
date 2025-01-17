package entities_test

import (
	"testing"

	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestNewURL(t *testing.T) {
	t.Run("Should create a new URL", func(t *testing.T) {
		// Arrange
		code := "asda12831"
		originalURL := "https://www.google.com"

		expected := entities.Url{
			OriginalURL: originalURL,
			Code:        code,
		}

		// Act
		url := entities.NewURL(originalURL, code)

		// Assert
		assert.Equal(t, expected.OriginalURL, url.OriginalURL)
		assert.Equal(t, expected.Code, url.Code)
		assert.NotEmpty(t, url.CreatedAt)
		assert.NotEmpty(t, url.UpdatedAt)
	})
}
