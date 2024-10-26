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

func (f ApiQueryParams) QueryString() string {
	v, _ := url.ParseQuery("")
	for key, value := range f {
		v.Set(key, fmt.Sprintf("%s", value))
	}
	return v.Encode()
}
