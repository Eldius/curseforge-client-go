package client

import (
	"fmt"
	"net/url"
)

type ApiQueryParams map[string]any

type ApiQueryOption func(ApiQueryParams)

// MinecraftVersionsQueryOption defines a ModLoader list query options
type MinecraftVersionsQueryOption ApiQueryOption

// WithSortDescending defines sort order to be descending
func WithSortDescending(descending bool) MinecraftVersionsQueryOption {
	return MinecraftVersionsQueryOption(func(m ApiQueryParams) {
		m["sortDescending"] = descending
	})
}

// MinecraftModLoadersQueryOption defines a ModLoader list query options
type MinecraftModLoadersQueryOption ApiQueryOption

// WithMinecraftVersion defines Minecraft for ModLoader search
func WithMinecraftVersion(version string) MinecraftModLoadersQueryOption {
	return MinecraftModLoadersQueryOption(func(m ApiQueryParams) {
		m["version"] = version
	})
}

// WithIncludeAll defines return a complete list or not
func WithIncludeAll(b bool) MinecraftModLoadersQueryOption {
	return MinecraftModLoadersQueryOption(func(m ApiQueryParams) {
		m["includeAll"] = b
	})
}

type ModsQueryOption ApiQueryOption

// WithGameID defines game ID for mod search
func WithGameID(g string) ModsQueryOption {
	return ModsQueryOption(func(m ApiQueryParams) {
		m["gameId"] = g
	})
}

// WithModsSeatchFilter defines query filter search
func WithModsSeatchFilter(q string) ModsQueryOption {
	return ModsQueryOption(func(m ApiQueryParams) {
		m["searchFilter"] = q
	})
}

// WithModsModLoaderType defines mod loader to be used
func WithModsModLoaderType(mlt string) ModsQueryOption {
	return ModsQueryOption(func(m ApiQueryParams) {
		m["modLoaderType"] = mlt
	})
}

// WithModsGameVersion defines game version to be used
func WithModsGameVersion(gv string) ModsQueryOption {
	return ModsQueryOption(func(m ApiQueryParams) {
		m["gameVersion"] = gv
	})
}

// WithModsGameID defines game ID to be used
func WithModsGameID(id string) ModsQueryOption {
	return ModsQueryOption(func(m ApiQueryParams) {
		m["gameId"] = id
	})
}

// WithModsClassID defines mod class ID
func WithModsClassID(cid string) ModsQueryOption {
	return ModsQueryOption(func(m ApiQueryParams) {
		m["classId"] = cid
	})
}

// WithModsCategoryID defines a category ID
func WithModsCategoryID(cid string) ModsQueryOption {
	return ModsQueryOption(func(m ApiQueryParams) {
		m["categoryId"] = cid
	})
}

func (f ApiQueryParams) QueryString() string {
	v, _ := url.ParseQuery("")
	for key, value := range f {
		v.Set(key, fmt.Sprintf("%s", value))
	}
	return v.Encode()
}