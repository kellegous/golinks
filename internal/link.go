package internal

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// Link represents a short link with a set of matchers that are responsible for using regex patterns to expand various URIs prefixed by the link's name.
type Link struct {
	Name    string    `json:"name"`
	Matches []*Match  `json:"matches"`
	Time    time.Time `json:"time"`
}

func (l *Link) UnmarshalJSON(data []byte) error {
	var t struct {
		Name    string    `json:"name"`
		Matches []*Match  `json:"matches"`
		Time    time.Time `json:"time"`
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	if strings.Contains(t.Name, "/") {
		return errors.New("prefix cannot contain '/'")
	}

	if len(t.Matches) == 0 {
		return errors.New("must have at least one match")
	}

	l.Name = t.Name
	l.Matches = t.Matches
	l.Time = t.Time
	return nil
}

func (l *Link) Expand(uri string) *ExpandedURL {
	if !strings.HasPrefix(uri, l.Name) {
		return nil
	}

	// remove the name prefix
	s := strings.TrimLeft(uri[len(l.Name):], "/")

	// find the first match that can expand the uri
	for i, match := range l.Matches {
		if expanded, ok := match.ExpandURL(s); ok {
			return &ExpandedURL{
				URL:   expanded,
				Index: i,
				Link:  l,
			}
		}
	}
	return nil
}

func (l *Link) Clone() *Link {
	matches := make([]*Match, len(l.Matches))
	for i, m := range l.Matches {
		matches[i] = m.Clone()
	}

	return &Link{
		Name:    l.Name,
		Matches: matches,
		Time:    l.Time,
	}
}

func LinksAreSame(a, b *Link) bool {
	if a == b {
		return true
	} else if a == nil || b == nil {
		return false
	}
	return a.Name == b.Name &&
		allMatchesAreSame(a.Matches, b.Matches) &&
		a.Time.Equal(b.Time)
}

func allMatchesAreSame(a, b []*Match) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !MatchesAreSame(a[i], b[i]) {
			return false
		}
	}
	return true
}
