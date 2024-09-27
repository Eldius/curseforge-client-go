package client

import (
	"github.com/eldius/curseforge-client-go/client/config"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/h2non/gock"
)

func TestGameList(t *testing.T) {
	defer gock.Off()
	gock.New(config.CurseforgeBaseURL).
		Get(gamesListPath).
		MatchHeader("x-api-key", "ABC123").
		Reply(200).
		File("samples/games_response_simplified.json")

	c := NewClient("ABC123")

	res, err := c.GetGames()
	assert.Nil(t, err)

	assert.Len(t, res.Data, 4)
	t.Log(res.Data)
}

func TestGetMods(t *testing.T) {
	defer gock.Off()
	gock.New(config.CurseforgeBaseURL).
		Get(modSearchPath).
		MatchHeader("x-api-key", "ABC123").
		MatchParam("gameId", "123").
		Reply(200).
		File("samples/minecraft_mods_category_modpacks_filter_atm_simplified.json")

	c := NewClient("ABC123")

	res, err := c.GetMods(ModFilter{GameID: "123"})
	assert.Nil(t, err)

	assert.Len(t, res.Data, 4)
	assert.NotEmpty(t, res.RawBody)
	t.Log(res.Data)
	t.Log(res.RawBody)
}

func TestGetModsByCategory(t *testing.T) {
	defer gock.Off()
	gock.New(config.CurseforgeBaseURL).
		Get(modSearchPath).
		MatchHeader("x-api-key", "ABC123").
		MatchParam("gameId", "123").
		MatchParam("categoryId", "modpacks").
		MatchParam("searchFilter", "atm").
		Reply(200).
		File("samples/minecraft_mods_category_modpacks_filter_atm_simplified.json")

	c := NewClient("ABC123")

	res, err := c.GetModsByCategory("123", "modpacks", "atm")
	assert.Nil(t, err)

	assert.Len(t, res.Data, 4)
	assert.NotEmpty(t, res.RawBody)

	t.Log(res.Data)
	t.Log(res.RawBody)

}
