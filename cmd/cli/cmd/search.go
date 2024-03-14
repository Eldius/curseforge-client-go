package cmd

import (
	"fmt"
	"github.com/eldius/curseforge-client-go/client"
	"github.com/eldius/curseforge-client-go/internal/config"
	"github.com/eldius/curseforge-client-go/internal/model"
	"github.com/eldius/curseforge-client-go/internal/persistence"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a Mod",
	Long:  `Search for a Mod.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		c := client.NewClient(config.GetCurseforgeAPIKey())
		_, _ = persistence.GetDB()
		res, err := c.GetMods("432")
		if err != nil {
			panic(err)
		}

		for _, m := range res.Data {
			fmt.Println(model.NewMod(m))
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
