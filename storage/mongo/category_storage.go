package mongo

import (
	"context"

	"github.com/mgajin/kpjm-scraper/common/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryStorage struct {
	collection *mongo.Collection
	storage    *Storage
}

func NewCategoryStorage(s *Storage) *CategoryStorage {
	collection := s.Collection("categories")

	return &CategoryStorage{
		collection: collection,
		storage:    s,
	}
}

func (cs *CategoryStorage) StoreCategory(category *model.Category) error {
	ctx := context.Background()
	_, err := cs.collection.InsertOne(ctx, category)

	return err
}
