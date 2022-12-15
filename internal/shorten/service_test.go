package shorten_test

import (
	"context"
	"testing"

	"github.com/dbashirov/link-shrinker/internal/model"
	"github.com/dbashirov/link-shrinker/internal/shorten"
	shortening "github.com/dbashirov/link-shrinker/internal/storage/shortering"
	"github.com/stretchr/testify/require"
)

func TestSerivice_Shortne(t *testing.T) {
	t.Run("generates shortening for a given URL", func(t *testing.T) {
		var (
			svc   = shorten.NewService(shortening.NewInMemory())
			input = model.ShortenInput{RawURL: "https://www.google.com"}
		)
		shortening, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)
	})
}
