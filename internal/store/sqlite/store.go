package sqlite

import (
	"context"
	"database/sql"
	"iter"

	"github.com/kellegous/golinks/internal"
	"github.com/kellegous/golinks/internal/store"
)

type Store struct {
	db *sql.DB
}

func (s *Store) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *Store) Put(
	ctx context.Context,
	link *internal.Link,
) error {
	return nil
}

func (s *Store) Get(
	ctx context.Context,
	name string,
) (*internal.Link, error) {
	return nil, nil
}

func (s *Store) Delete(
	ctx context.Context,
	name string,
) error {
	return nil
}

func (s *Store) List(
	ctx context.Context,
	opts *store.ListOptions,
) iter.Seq2[*internal.Link, error] {
	return nil
}

func FromDSN(ctx context.Context, dsn string) (*Store, error) {
	return &Store{}, nil
}
