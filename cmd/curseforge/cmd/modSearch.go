package cmd

import (
	"context"
	"fmt"
	"github.com/eldius/curseforge-client-go/client"
	client_config "github.com/eldius/curseforge-client-go/client/config"
	"github.com/eldius/curseforge-client-go/internal/config"
	"github.com/eldius/curseforge-client-go/internal/model"
	"github.com/eldius/curseforge-client-go/internal/persistence"
	"github.com/spf13/cobra"
	"log/slog"
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
		c.SetLogger(client.NewDefaultSlogClientLogger(slog.Default()))

		index := int64(0)

		ctx := context.Background()
		for {
			mods, err := c.GetMods(client.ModFilter{
				GameID:      modSearchOptions.gameID,
				Term:        modSearchOptions.searchTerm,
				GameVersion: modSearchOptions.gameVersion,
				ClassID:     modSearchOptions.classID,
				Index:       index,
				PageSize:    50,
			})
			if err != nil {
				err = fmt.Errorf("looking for mods: %w", err)
				panic(err)
			}

			slog.With(
				slog.Bool("fetch_all", modSearchOptions.fetchAll),
				slog.Int("result_count", mods.Pagination.ResultCount),
				slog.Int("total_count", mods.Pagination.TotalCount),
				slog.Int("page_size", mods.Pagination.PageSize),
				slog.Int("pagination_index", mods.Pagination.Index),
				slog.Int64("index", index),
			).Debug("modSearch")

			if !modSearchOptions.fetchAll || (mods.Pagination.ResultCount <= 0) {
				break
			}

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
			} else {
				fmt.Println("found", mods.Pagination.ResultCount, "mods for this search...")
				for _, m := range mods.Data {
					fmt.Println(m.String())
				}
			}
			index += int64(mods.Pagination.PageSize)
		}
	},
}

var (
	modSearchOptions struct {
		searchTerm  string
		gameVersion string
		gameID      string
		classID     string
		saveCache   bool
		fetchAll    bool
	}
)

func init() {
	modCmd.AddCommand(modSearchCmd)

	modSearchCmd.Flags().StringVar(&modSearchOptions.searchTerm, "term", "", "Search term to use")
	modSearchCmd.Flags().StringVar(&modSearchOptions.gameVersion, "game-version", "", "Game version")
	modSearchCmd.Flags().StringVar(&modSearchOptions.gameID, "game-id", "", "Game ID (from game search)")
	modSearchCmd.Flags().StringVar(&modSearchOptions.classID, "class-id", "", "Mod class ID (discoverable via Categories)")
	modSearchCmd.Flags().BoolVar(&modSearchOptions.saveCache, "save-cache", false, "Save to a local cache")
	modSearchCmd.Flags().BoolVar(&modSearchOptions.fetchAll, "fetch-all", false, "Still fetching results until hit the end of the list")
}
