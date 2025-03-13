package types

import "strings"

// ModLoaderByName returns a modloader reference from name
func ModLoaderByName(ml string) ModLoaderType {
	switch strings.ToLower(ml) {
	case "forge":
		return ModLoaderTypeForge
	case "cauldron":
		return ModLoaderTypeCauldron
	case "liteloader":
		return ModLoaderTypeLiteLoader
	case "fabric":
		return ModLoaderTypeFabric
	case "quilt":
		return ModLoaderTypeQuilt
	case "neoforge":
		return ModLoaderTypeNeoForge
	default:
		return ModLoaderTypeAny
	}
}
