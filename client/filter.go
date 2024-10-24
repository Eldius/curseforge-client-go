package client

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrorModFilterIsEmpty = errors.New("mod filter is empty")
)

type ModLoaderType int

type ModStatus int

const (
	ModLoaderTypeAny        ModLoaderType = 0
	ModLoaderTypeForge      ModLoaderType = 1
	ModLoaderTypeCauldron   ModLoaderType = 2
	ModLoaderTypeLiteLoader ModLoaderType = 3
	ModLoaderTypeFabric     ModLoaderType = 4
	ModLoaderTypeQuilt      ModLoaderType = 5
	ModLoaderTypeNeoForge   ModLoaderType = 6
)

const (
	ModStatusNew             ModStatus = 1
	ModStatusChangesRequired ModStatus = 2
	ModStatusUnderSoftReview ModStatus = 3
	ModStatusApproved        ModStatus = 4
	ModStatusRejected        ModStatus = 5
	ModStatusChangesMade     ModStatus = 6
	ModStatusInactive        ModStatus = 7
	ModStatusAbandoned       ModStatus = 8
	ModStatusDeleted         ModStatus = 9
	ModStatusUnderReview     ModStatus = 10
)

type ModFilter struct {
	GameID        string
	Term          string
	GameVersion   string
	ClassID       string
	ModLoaderType ModLoaderType
	ModStatus     ModStatus
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
