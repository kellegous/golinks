package memory

import (
	"context"
	"iter"
	"sync"

	"github.com/kellegous/golinks/internal"
	"github.com/kellegous/golinks/internal/store"
	"github.com/tidwall/btree"
)

type Store struct {
	lck   sync.RWMutex
	links btree.Map[string, *internal.Link]
}

func (s *Store) Close(ctx context.Context) error {
	return nil
}

func (s *Store) Put(
	ctx context.Context,
	l *internal.Link,
) error {
	s.lck.Lock()
	defer s.lck.Unlock()

	s.links.Set(l.Name, l.Clone())
	return nil
}

func (s *Store) Get(
	ctx context.Context,
	name string,
) (*internal.Link, error) {
	s.lck.RLock()
	defer s.lck.RUnlock()

	if v, ok := s.links.Get(name); ok {
		return v, nil
	}
	return nil, store.ErrNotFound
}

func (s *Store) Delete(
	ctx context.Context,
	name string,
) error {
	s.lck.Lock()
	defer s.lck.Unlock()

	if _, ok := s.links.Delete(name); !ok {
		return store.ErrNotFound
	}
	return nil
}

func (s *Store) List(
	ctx context.Context,
	opts interface{},
) iter.Seq2[*internal.Link, error] {
	return func(yield func(*internal.Link, error) bool) {
		s.lck.RLock()
		defer s.lck.RUnlock()

		s.links.Scan(func(k string, v *internal.Link) bool {
			return yield(v, nil)
		})
	}
}
