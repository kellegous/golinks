package store

import (
	"context"
	"iter"

	"github.com/kellegous/golinks/internal"
)

type Store interface {
	Close(ctx context.Context) error
	Put(ctx context.Context, l *internal.Link) error
	Get(ctx context.Context, name string) (*internal.Link, error)
	List(ctx context.Context, opts any) iter.Seq2[*internal.Link, error]
	Delete(ctx context.Context, name string) error
}
