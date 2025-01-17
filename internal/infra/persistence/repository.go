package persistence

import (
	"context"

	"github.com/jfelipearaujo-urlshortener/urlshortener-app/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName string = "urls"
)

type repository struct {
	collection *mongo.Collection
}

func New(db *mongo.Database) UrlRepository {
	return &repository{
		collection: db.Collection(collectionName),
	}
}

func (rep *repository) Create(ctx context.Context, url *entities.Url) error {
	result, err := rep.collection.InsertOne(ctx, url)
	if err != nil {
		return err
	}

	url.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}
