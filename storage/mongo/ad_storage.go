package mongo

import (
	"context"

	"github.com/mgajin/kpjm-scraper/common/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdStorage struct {
	collection *mongo.Collection
	storage    *Storage
}

func NewAdStorage(s *Storage) *AdStorage {
	collection := s.Collection("ads")

	return &AdStorage{
		collection: collection,
		storage:    s,
	}
}

func (as *AdStorage) StoreAd(ad *model.Ad) error {
	ctx := context.Background()
	_, err := as.collection.InsertOne(ctx, ad)

	return err
}
