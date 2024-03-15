package logger

import (
	"fmt"
	"github.com/eldius/curseforge-client-go/client"
	"log/slog"
)

type ClientLogger struct {
	client.Logger
}

func (l ClientLogger) Printf(format string, args ...any) {
	slog.Debug(fmt.Sprintf(format, args))
}

func (l ClientLogger) Println(v ...any) {
	slog.Debug(fmt.Sprintf("%v", v))
}
