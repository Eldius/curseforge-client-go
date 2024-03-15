package client

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrorModFilterIsEmpty = errors.New("mod filter is empty")
)

type ModFilter struct {
	GameID      string
	Term        string
	GameVersion string
}

func (f ModFilter) QueryString() string {
	v, _ := url.ParseQuery("")
	if strings.TrimSpace(f.GameID) != "" {
		v.Set("gameId", f.GameID)
	}
	if strings.TrimSpace(f.Term) != "" {
		v.Set("searchFilter", f.Term)
	}
	if strings.TrimSpace(f.GameVersion) != "" {
		v.Set("gameVersion", f.GameVersion)
	}
	return v.Encode()
}
