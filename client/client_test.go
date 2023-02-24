package client

import (
	"github.com/eldius/curseforge-client-go/client/config"
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
	if err != nil {
		t.Log("Failed to list games")
		t.FailNow()
	}

	if len(res.Data) != 4 {
		t.Logf("Should return 4 games, but returned %d", len(res.Data))
		t.FailNow()
	}
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

	res, err := c.GetMods("123")
	if err != nil {
		t.Log("Failed to list games")
		t.FailNow()
	}

	if len(res.Data) != 4 {
		t.Logf("Should return 4 games, but returned %d", len(res.Data))
		t.FailNow()
	}
	t.Log(res.Data)

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
	if err != nil {
		t.Log("Failed to list games")
		t.FailNow()
	}

	if len(res.Data) != 4 {
		t.Logf("Should return 4 games, but returned %d", len(res.Data))
		t.FailNow()
	}
	t.Log(res.Data)

}
