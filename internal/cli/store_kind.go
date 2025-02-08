package cli

import "fmt"

type StoreKind string

const (
	StoreTypeMemory  StoreKind = "memory"
	StoreTypeSQLite  StoreKind = "sqlite"
	StoreTypeLevelDB StoreKind = "leveldb"
)

func kindFromString(v string) (StoreKind, error) {
	switch v {
	case "memory", "mem":
		return StoreTypeMemory, nil
	case "sql":
		return StoreTypeSQLite, nil
	case "leveldb":
		return StoreTypeLevelDB, nil
	default:
		return "", fmt.Errorf("unknown store type %q", v)
	}
}
