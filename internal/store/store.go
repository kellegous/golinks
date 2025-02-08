package store

import (
	"context"
	"errors"
	"iter"

	"github.com/kellegous/golinks/internal"
)

var ErrNotFound = errors.New("not found")

type Store interface {
	Close(ctx context.Context) error
	Put(ctx context.Context, l *internal.Link) error
	Get(ctx context.Context, name string) (*internal.Link, error)
	List(ctx context.Context, opts *ListOptions) iter.Seq2[*internal.Link, error]
	Delete(ctx context.Context, name string) error
}
