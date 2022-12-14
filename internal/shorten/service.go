package shorten

import (
	"context"

	"github.com/dbashirov/link-shrinker/internal/model"
	"github.com/google/uuid"
)

type Storage interface {
	Put(ctx context.Context, shortening model.Shortening) (*model.Shortening, error)
	Get(ctx context.Context, identifier string) (*model.Shortening, error)
	IncrementVisits(ctx context.Context, identifier string) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Shorten(ctx context.Context, input model.ShortenInput) (*model.Shortening, error) {
	// 1. Сгененировать сокращенный идентификатор
	var (
		id         = uuid.New().ID()
		identifier string
	)

	if input.Identifier == "" {
		identifier = Shorten(id)
	} else {
		identifier = input.Identifier
	}

	// 2. Сохранить в хранилище
	dbShortening := model.Shortening{
		Identifier:  identifier,
		OriginalURL: input.RawURL,
	}

	shortening, err := s.storage.Put(ctx, dbShortening)
	if err != nil {
		return nil, err
	}

	return shortening, nil
}
