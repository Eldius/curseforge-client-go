package cli_config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// Setup configures app parameters
func Setup(cfgFile string) error {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mock-server-go" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".bign-price-fetcher"))
		viper.AddConfigPath(filepath.Join(home))
		viper.SetConfigType("yaml")
		viper.SetConfigName("bign")
	}

	SetDefaults()
	MapEnvVars()

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
		return err
	}

	if err := setupLogs(); err != nil {
		err = fmt.Errorf("failed to configure logs: %w", err)
		return err
	}

	return nil
}

// SetDefaults sets default configuration values
func SetDefaults() {
	viper.SetDefault(LogFormatKey, "text")
	viper.SetDefault(LogLevelKey, LogLevelINFO)
}

// MapEnvVars sets up environment variables mapping
func MapEnvVars() {
	viper.SetEnvPrefix("curseforge")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func setupLogs() error {
	var h slog.Handler
	var w io.Writer = os.Stdout
	if out := GetLogOutput(); out != "" {
		out, err := filepath.Abs(out)
		if err != nil {
			err = fmt.Errorf("parsing log absolute path: %w", err)
			return err
		}
		f, err := os.Create(out)
		if err != nil {
			err = fmt.Errorf("opening log file: %w", err)
			return err
		}
		w = io.MultiWriter(w, f)
	}
	if GetLogFormat() == LogFormatJSON {
		h = slog.NewJSONHandler(w, &slog.HandlerOptions{
			AddSource: true,
			Level:     parseLogLevel(GetLogLevel()),
		})
	} else {
		h = slog.NewTextHandler(w, &slog.HandlerOptions{
			AddSource: true,
			Level:     parseLogLevel(GetLogLevel()),
		})
	}
	logger := slog.New(h)
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	slog.SetDefault(logger.With(
		slog.String("service", "bign-fetcher"),
		slog.String("host", host),
	))

	return nil
}

func parseLogLevel(lvl string) slog.Level {
	switch lvl {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
