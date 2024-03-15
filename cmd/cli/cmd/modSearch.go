package cmd

import (
	"context"
	"fmt"
	"github.com/eldius/curseforge-client-go/client"
	client_config "github.com/eldius/curseforge-client-go/client/config"
	"github.com/eldius/curseforge-client-go/internal/config"
	"github.com/eldius/curseforge-client-go/internal/logger"
	"github.com/eldius/curseforge-client-go/internal/model"
	"github.com/eldius/curseforge-client-go/internal/persistence"
	"github.com/spf13/cobra"
)

// modSearchCmd represents the search command
var modSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a Mod",
	Long:  `Search for a Mod.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClientWithConfig(client_config.NewConfig(
			config.GetCurseforgeAPIKey(),
			client_config.EnableDebug(true),
		))
		c.SetLogger(logger.SlogClientLogger{})
		mods, err := c.GetMods(client.ModFilter{
			GameID:      modSearchOptions.gameID,
			Term:        modSearchOptions.searchTerm,
			GameVersion: modSearchOptions.gameVersion,
		})
		if err != nil {
			err = fmt.Errorf("looking for mods: %w", err)
			panic(err)
		}

		ctx := context.Background()
		if modSearchOptions.saveCache {
			db, err := persistence.GetDB()
			if err != nil {
				panic(err)
			}
			r := persistence.NewModRepository(db)
			for _, m := range mods.Data {
				r.Save(ctx, model.NewMod(m))
				fmt.Println(m.String())
			}
			return
		}
		fmt.Println("found", mods.Pagination.ResultCount, "mods for this search...")
		for _, m := range mods.Data {
			fmt.Println(m.String())
		}
	},
}

var (
	modSearchOptions struct {
		searchTerm  string
		gameVersion string
		gameID      string
		saveCache   bool
	}
)

func init() {
	modCmd.AddCommand(modSearchCmd)

	modSearchCmd.Flags().StringVar(&modSearchOptions.searchTerm, "term", "", "Search term to use")
	modSearchCmd.Flags().StringVar(&modSearchOptions.gameVersion, "game-version", "", "Game version")
	modSearchCmd.Flags().StringVar(&modSearchOptions.gameID, "game-id", "", "Game ID (from game search)")
	modSearchCmd.Flags().BoolVar(&modSearchOptions.saveCache, "save-cache", false, "Save to a local cache")
}
