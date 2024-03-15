package cmd

import (
	"fmt"
	"github.com/eldius/curseforge-client-go/client"
	client_config "github.com/eldius/curseforge-client-go/client/config"
	"github.com/eldius/curseforge-client-go/internal/config"
	"github.com/eldius/curseforge-client-go/internal/logger"
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
			GameID:      searchOptions.gameID,
			Term:        searchOptions.searchTerm,
			GameVersion: searchOptions.gameVersion,
		})
		if err != nil {
			err = fmt.Errorf("looking for mods: %w", err)
			panic(err)
		}

		fmt.Println("found", mods.Pagination.ResultCount, "mods for this search...")
		for _, m := range mods.Data {
			fmt.Println(m.String())
		}
	},
}

var (
	searchOptions struct {
		searchTerm  string
		gameVersion string
		gameID      string
	}
)

func init() {
	modCmd.AddCommand(modSearchCmd)

	modSearchCmd.Flags().StringVar(&searchOptions.searchTerm, "term", "", "Search term to use")
	modSearchCmd.Flags().StringVar(&searchOptions.gameVersion, "game-version", "", "Game version")
	modSearchCmd.Flags().StringVar(&searchOptions.gameID, "game-id", "", "Game ID (from game search)")
}
