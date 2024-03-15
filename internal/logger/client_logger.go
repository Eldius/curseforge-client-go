package logger

import (
	"bytes"
	"fmt"
	"github.com/eldius/curseforge-client-go/client"
	"io"
	"log/slog"
	"net/http"
)

type SlogClientLogger struct {
	client.Logger
}

func (l SlogClientLogger) Printf(format string, args ...any) {
	slog.Debug(fmt.Sprintf(format, args))
}

func (l SlogClientLogger) Println(v ...any) {
	slog.Debug(fmt.Sprintf("%v", v))
}

func (l SlogClientLogger) DebugRequest(res *http.Response) {
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
