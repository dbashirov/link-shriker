package shorten_test

import (
	"testing"

	"github.com/dbashirov/link-shrinker/internal/shorten"
	shortening "github.com/dbashirov/link-shrinker/internal/storage/shortering"
)

func TestSerivice_Shortne(t *testing.T) {
	t.Run("generates shortening for a given URL", func(t *testing.T) {
		var (
			svc = shorten.NewService(shortening.NewInMemory())
		)
	})
}
