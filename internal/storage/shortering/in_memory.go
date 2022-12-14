package shortening

import (
	"context"
	"sync"
	"time"

	"github.com/dbashirov/link-shrinker/internal/model"
)

type inMemory struct {
	m sync.Map
}

func NewInMemory() *inMemory {
	return &inMemory{}
}

func (s *inMemory) Put(ctx context.Context, shortening model.Shortening) (*model.Shortening, error) {

	if _, exists := s.m.Load(shortening.Identifier); exists {
		return nil, model.ErrIdentifierExists
	}

	shortening.CreatedAt = time.Now().UTC()

	s.m.Store(shortening.Identifier, shortening)

	return &shortening, nil
}

func (s *inMemory) Get(ctx context.Context, identifier string) (*model.Shortening, error) {
	panic("implement me")
}

func (s *inMemory) IncrementVisits(ctx context.Context, identifier string) error {
	panic("implement me")
}
