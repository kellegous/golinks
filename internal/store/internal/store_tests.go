package internal

import (
	"context"
	"errors"
	"iter"
	"regexp"
	"testing"
	"time"

	"github.com/kellegous/golinks/internal"
	"github.com/kellegous/golinks/internal/store"
)

func toSlice[T any](iter iter.Seq2[T, error]) ([]T, error) {
	var res []T

	for v, err := range iter {
		if err != nil {
			return nil, err
		}

		res = append(res, v)
	}

	return res, nil
}

func allLinksAreSame(a, b []*internal.Link) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !internal.LinksAreSame(a[i], b[i]) {
			return false
		}
	}

	return true
}

func TestStore(t *testing.T, s store.Store) {
	t.Run("GetPut", func(t *testing.T) {
		a := &internal.Link{
			Name: "a",
			Matches: []*internal.Match{
				{
					URIPattern:  regexp.MustCompile("^/a/(.*)$"),
					URLTemplate: "https://a.com/$1",
				},
			},
			Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		if _, err := s.Get(context.Background(), "a"); !errors.Is(err, store.ErrNotFound) {
			t.Fatalf("expected not found, got %v", err)
		}

		if err := s.Put(context.Background(), a); err != nil {
			t.Fatal(err)
		}

		b, err := s.Get(context.Background(), "a")
		if err != nil {
			t.Fatal(err)
		}

		if !internal.LinksAreSame(a, b) {
			t.Fatalf("got %#v, expected %#v", b, a)
		}

		if a == b {
			t.Fatal("expected clone")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		a := &internal.Link{
			Name: "a",
			Matches: []*internal.Match{
				{
					URIPattern:  regexp.MustCompile("^/a/(.*)$"),
					URLTemplate: "https://a.com/$1",
				},
			},
			Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		if err := s.Put(context.Background(), a); err != nil {
			t.Fatal(err)
		}

		if err := s.Delete(context.Background(), "a"); err != nil {
			t.Fatal(err)
		}

		if _, err := s.Get(context.Background(), "a"); !errors.Is(err, store.ErrNotFound) {
			t.Fatalf("expected not found, got %v", err)
		}
	})

	t.Run("List", func(t *testing.T) {
		a := &internal.Link{
			Name: "a",
			Matches: []*internal.Match{
				{
					URIPattern:  regexp.MustCompile("^/a/(.*)$"),
					URLTemplate: "https://a.com/$1",
				},
			},
			Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		b := &internal.Link{
			Name: "b",
			Matches: []*internal.Match{
				{
					URIPattern:  regexp.MustCompile("^/b/(.*)$"),
					URLTemplate: "https://b.com/$1",
				},
			},
			Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		c := &internal.Link{
			Name: "c",
			Matches: []*internal.Match{
				{
					URIPattern:  regexp.MustCompile("^/c/(.*)$"),
					URLTemplate: "https://c.com/$1",
				},
			},
			Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		if err := s.Put(context.Background(), b); err != nil {
			t.Fatal(err)
		}

		{
			all, err := toSlice(s.List(context.Background(), nil))
			if err != nil {
				t.Fatal(err)
			}

			if !allLinksAreSame(all, []*internal.Link{b}) {
				t.Fatalf("expected %#v, got %#v", []*internal.Link{b}, all)
			}
		}

		if err := s.Put(context.Background(), a); err != nil {
			t.Fatal(err)
		}

		{
			all, err := toSlice(s.List(context.Background(), nil))
			if err != nil {
				t.Fatal(err)
			}

			if !allLinksAreSame(all, []*internal.Link{a, b}) {
				t.Fatalf("expected %#v, got %#v", []*internal.Link{b}, all)
			}
		}

		if err := s.Put(context.Background(), c); err != nil {
			t.Fatal(err)
		}

		{
			all, err := toSlice(s.List(context.Background(), nil))
			if err != nil {
				t.Fatal(err)
			}

			if !allLinksAreSame(all, []*internal.Link{a, b, c}) {
				t.Fatalf("expected %#v, got %#v", []*internal.Link{b}, all)
			}
		}
	})
}
