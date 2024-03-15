package cmd

import (
	"fmt"
	"github.com/eldius/curseforge-client-go/internal/config"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "curseforge-client-go",
	Short: "A simple CLI to interact with CurseForge API",
	Long:  `A simple CLI to interact with CurseForge API.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return config.Setup(cfgFile)
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return config.Setup(cfgFile)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	cfgFile string
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.curseforge-client-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug log")
	if err := viper.BindPFlag(config.DebugEnabled, rootCmd.PersistentFlags().Lookup("debug")); err != nil {
		err = fmt.Errorf("binding debug parameter: %w", err)
		panic(err)
	}
}
