package shorten_test

import (
	"context"
	"testing"

	"github.com/dbashirov/link-shrinker/internal/model"
	"github.com/dbashirov/link-shrinker/internal/shorten"
	shortening "github.com/dbashirov/link-shrinker/internal/storage/shortering"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerivice_Shorten(t *testing.T) {
	t.Run("generates shortening for a given URL", func(t *testing.T) {
		var (
			svc   = shorten.NewService(shortening.NewInMemory())
			input = model.ShortenInput{RawURL: "https://www.google.com"}
		)
		shortening, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)

		assert.NotEmpty(t, shortening.Identifier)
		assert.Equal(t, "https://www.google.com", shortening.OriginalURL)
		assert.NotZero(t, shortening.CreatedAt)
	})

	t.Run("uses custom identifier if provided", func(t *testing.T) {
		const identifier = "google"
		var (
			svc   = shorten.NewService(shortening.NewInMemory())
			input = model.ShortenInput{
				RawURL:     "https://www.google.com",
				Identifier: identifier,
			}
		)
		shortening, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)

		assert.Equal(t, identifier, shortening.Identifier)
		assert.Equal(t, "https://www.google.com", shortening.OriginalURL)
		assert.NotZero(t, shortening.CreatedAt)
	})

	t.Run("returns error if identifier is already taken", func(t *testing.T) {
		const identifier = "google"
		var ()
	})
}
