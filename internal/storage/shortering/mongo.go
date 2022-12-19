package shortening

import (
	"context"
	"fmt"
	"time"

	"github.com/dbashirov/link-shrinker/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mgo struct {
	db *mongo.Database
}

func NewMongoDB(client *mongo.Client) *mgo {
	return &mgo{db: client.Database("url-shortener")}
}

func (m *mgo) col() *mongo.Collection {
	return m.db.Collection("shortenings")
}

func (m *mgo) Put(ctx context.Context, shortening model.Shortening) (*model.Shortening, error) {
	const op = "shortening.mgo.Put"

	shortening.CreatedAt = time.Now().UTC()

	// 1. Проверка, нет ли в коллекции документа с таким же идентификатором
	count, err := m.col().CountDocuments(ctx, bson.M{"_id": shortening.Identifier})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if count > 0 {
		return nil, fmt.Errorf("%s: %w", op, model.ErrIdentifierExists)
	}

	// 2. Если нет, то добавить документ в коллекцию
	_, err = m.col().InsertOne(ctx, mgoShorteningFromModel(shortening))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &shortening, nil
}

func (s *mgo) Get(ctx context.Context, identifier string) (*model.Shortening, error) {
	var op = "shortening.mgo.get"

	var shortening mgoShortening
	if err := s.col().FindOne(ctx, bson.M{"_id": identifier}).Decode(&shortening); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%s: %w", op, model.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return modelShorteningFromMgo(shortening), nil
}

func (s *mgo) IncrementVisits(ctx context.Context, identifier string) error {
	panic("implement me")
}

type mgoShortening struct {
	Identifier  string    `bson:"_id"`
	CreatedBy   string    `bson:"created_by"`
	OriginalURL string    `bson:"original_url"`
	Visits      int64     `bson:"visits"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func mgoShorteningFromModel(shortening model.Shortening) mgoShortening {
	return mgoShortening{
		Identifier:  shortening.Identifier,
		CreatedBy:   shortening.CreatedBy,
		OriginalURL: shortening.OriginalURL,
		Visits:      shortening.Visits,
		CreatedAt:   shortening.CreatedAt,
		UpdatedAt:   shortening.UpdatedAt,
	}
}

func modelShorteningFromMgo(shortening mgoShortening) *model.Shortening {
	return &model.Shortening{
		Identifier:  shortening.Identifier,
		CreatedBy:   shortening.CreatedBy,
		OriginalURL: shortening.OriginalURL,
		Visits:      shortening.Visits,
		CreatedAt:   shortening.CreatedAt,
		UpdatedAt:   shortening.UpdatedAt,
	}
}
