package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/kellegous/golinks/internal/store"
	"github.com/kellegous/golinks/internal/store/memory"
	"github.com/kellegous/golinks/internal/store/sqlite"
)

type StoreConfig struct {
	Kind StoreKind
	DSN  string
}

func (c *StoreConfig) Set(s string) error {
	var err error
	kind, dsn, found := strings.Cut(s, ":")
	if !found {
		kind = s
		dsn = ""
	}

	c.DSN = dsn
	c.Kind, err = kindFromString(kind)
	return err
}

func (c *StoreConfig) String() string {
	return ""
}

func (c *StoreConfig) Type() string {
	return "store"
}

func (c *StoreConfig) openStore(
	ctx context.Context,
) (store.Store, error) {
	switch c.Kind {
	case StoreTypeMemory:
		return memory.FromDSN(c.DSN)
	case StoreTypeSQLite:
		return sqlite.FromDSN(ctx, c.DSN)
	}
	return nil, fmt.Errorf("unknown store type %q", c.Kind)
}
