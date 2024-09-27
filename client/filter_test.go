package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModFilter(t *testing.T) {
	t.Run("given an empty mod filter should return an empty query string", func(t *testing.T) {
		f := ModFilter{}

		q := f.QueryString()
		assert.Equal(t, "index=0", q)
	})

	t.Run("given a mod filter with game id should return a query string with only game id parameter", func(t *testing.T) {
		f := ModFilter{
			GameID: "123",
		}

		q := f.QueryString()
		assert.NotEmpty(t, q)
		assert.Contains(t, q, "gameId")
		assert.Contains(t, q, f.GameID)
	})

	t.Run("given a mod filter with game id and game version should return a query string with both parameters", func(t *testing.T) {
		f := ModFilter{
			GameID:      "123",
			GameVersion: "456",
		}

		q := f.QueryString()
		assert.NotEmpty(t, q)
		assert.Contains(t, q, "gameId")
		assert.Contains(t, q, f.GameID)
		assert.Contains(t, q, "gameVersion")
		assert.Contains(t, q, f.GameVersion)
		assert.NotContains(t, q, "searchFilter")
	})

	t.Run("given a mod filter with game id and term should return a query string with both parameters", func(t *testing.T) {
		f := ModFilter{
			GameID: "123",
			Term:   "term",
		}

		q := f.QueryString()
		assert.NotEmpty(t, q)
		assert.Contains(t, q, "gameId")
		assert.Contains(t, q, f.GameID)
		assert.Contains(t, q, "searchFilter")
		assert.Contains(t, q, f.Term)
		assert.NotContains(t, q, "gameVersion")
	})
}
