package internal

import "time"

// Link represents a short link with a set of matchers that are responsible for using regex patterns to expand various URIs prefixed by the link's name.
type Link struct {
	Name    string    `json:"name"`
	Matches []*Match  `json:"matches"`
	Time    time.Time `json:"time"`
}
