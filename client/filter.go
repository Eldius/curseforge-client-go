package client

import (
	"errors"
	"github.com/eldius/curseforge-client-go/client/types"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrorModFilterIsEmpty = errors.New("mod filter is empty")
)

type ModFilter struct {
	GameID        string
	Term          string
	GameVersion   string
	ClassID       string
	ModLoaderType types.ModLoaderType
	ModStatus     types.ModStatus
	PageSize      int64
	Index         int64
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
	if strings.TrimSpace(f.ClassID) != "" {
		v.Set("classId", f.ClassID)
	}
	if f.PageSize != 0 {
		v.Set("pageSize", strconv.FormatInt(f.PageSize, 10))
	}

	if f.ModStatus != 0 {
		v.Set("pageSize", strconv.FormatInt(f.PageSize, 10))
	}

	v.Set("modLoaderType", strconv.Itoa(int(f.ModLoaderType)))
	v.Set("index", strconv.FormatInt(f.Index, 10))
	return v.Encode()
}
