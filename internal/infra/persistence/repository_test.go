package persistence_test

import (
	"context"
	"testing"

	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/domain/entities"
	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/infra/persistence"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewUrlRepository(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should create a new UrlRepository", func(mt *mtest.T) {
		// Arrange

		// Act
		rep := persistence.New(mt.DB)

		// Assert
		assert.NotNil(mt, rep)
	})
}

func TestCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should create a new URL", func(mt *mtest.T) {
		// Arrange
		ctx := context.Background()

		sut := persistence.New(mt.DB)

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		url := entities.NewURL("https://www.google.com", "asda12831")

		// Act
		err := sut.Create(ctx, url)

		// Assert
		assert.NoError(mt, err)
	})

	mt.Run("Should return an error when trying to create a new URL", func(mt *mtest.T) {
		// Arrange
		ctx := context.Background()

		sut := persistence.New(mt.DB)

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: -1},
		})

		url := entities.NewURL("https://www.google.com", "asda12831")

		// Act
		err := sut.Create(ctx, url)

		// Assert
		assert.Error(mt, err)
	})
}
