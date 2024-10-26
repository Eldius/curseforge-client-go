package client

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type DefaultSlogClientLogger struct {
	Logger
	l *slog.Logger
}

func NewDefaultSlogClientLogger(logger *slog.Logger) Logger {
	return &DefaultSlogClientLogger{l: logger}
}

func (l DefaultSlogClientLogger) Printf(format string, args ...any) {
	slog.Debug(fmt.Sprintf(format, args...))
}

func (l DefaultSlogClientLogger) Println(v ...any) {
	slog.Debug(fmt.Sprintf("%v", v))
}

func (l DefaultSlogClientLogger) DebugRequest(res *http.Response) {
	buff, _ := io.ReadAll(res.Body)
	res.Body = io.NopCloser(bytes.NewReader(buff))
	slog.With(
		slog.String("url", res.Request.URL.String()),
		slog.String("req_uri", res.Request.RequestURI),
		slog.String("res", string(buff)),
		slog.String("status", res.Status),
		slog.Int("status_code", res.StatusCode),
	).Debug("CurseforgeAPIRequest")
}
