package mongodb

import (
	"context"
	"fmt"

	"github.com/snufkin23/web-crawler.git/internal/domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	DatabaseName   = "webcrawler_db"
	CollectionName = "pages"
)

type MongoArchiver struct {
	collection *mongo.Collection
}

func NewMongoArchiver(collection *mongo.Collection) *MongoArchiver {
	return &MongoArchiver{
		collection: collection,
	}
}

func (m *MongoArchiver) Save(ctx context.Context, page *domain.Page) (bool, error) {

	if _, err := m.collection.InsertOne(ctx, page); err != nil {
		return false, fmt.Errorf("failed to insert page into mongodb: %w", err)
	}
	return true, nil

}
