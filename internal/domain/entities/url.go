package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`

	OriginalURL string `json:"original_url" bson:"original_url"`
	Code        string `json:"code" bson:"code"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewURL(originalURL string, code string) *Url {
	return &Url{
		OriginalURL: originalURL,
		Code:        code,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
