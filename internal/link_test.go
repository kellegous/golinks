package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"
)

func TestLinkUnmarshalJSON(t *testing.T) {
	tests := []struct {
		JSON          string
		ExpectedLink  *Link
		ExpectedError error
	}{
		{
			`{
				"name":"a",
				"matches":[
					{"uri_pattern":"^/a/(.*)$","url_template":"https://a.com/$1"},
					{"uri_pattern":"^/b/(.*)$","url_template":"https://b.com/$1"}
				],
				"time":"2020-01-01T00:00:00Z"
			}`,
			&Link{
				Name: "a",
				Matches: []*Match{
					{
						URIPattern:  regexp.MustCompile("^/a/(.*)$"),
						URLTemplate: "https://a.com/$1",
					},
					{
						URIPattern:  regexp.MustCompile("^/b/(.*)$"),
						URLTemplate: "https://b.com/$1",
					},
				},
				Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			nil,
		},
		{
			`{
				"name":"a/b",
				"matches":[
					{"uri_pattern":"^/a/(.*)$","url_template":"https://a.com/$1"}
				],
				"time":"2020-01-01T00:00:00Z"
			}`,
			nil,
			errors.New("prefix cannot contain '/'"),
		},
		{
			`{
				"name":"a",
				"matches":[],
				"time":"2020-01-01T00:00:00Z"
			}`,
			nil,
			errors.New("must have at least one match"),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test-%d", i), func(t *testing.T) {
			var l Link
			if err := json.Unmarshal([]byte(test.JSON), &l); !errorsAreSame(err, test.ExpectedError) {
				t.Errorf("unexpected error: %v", err)
			}
			if test.ExpectedLink != nil && !LinksAreSame(&l, test.ExpectedLink) {
				t.Fatalf("for %s got %v expected %v", test.JSON, l, test.ExpectedLink)
			}
		})
	}
}

func expandedURLsAreSame(a, b *ExpandedURL) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.URL == b.URL && a.Index == b.Index && LinksAreSame(a.Link, b.Link)
}

func TestExpand(t *testing.T) {
	linkA := Link{
		Name: "a",
		Matches: []*Match{
			{
				URIPattern:  regexp.MustCompile("^(.*)$"),
				URLTemplate: "https://a.com/$1",
			},
		},
	}

	linkB := Link{
		Name: "b",
		Matches: []*Match{
			{
				URIPattern:  regexp.MustCompile("(^foo$)"),
				URLTemplate: "https://b.com/$1",
			},
			{
				URIPattern:  regexp.MustCompile("(^bar$)"),
				URLTemplate: "https://b.com/$1",
			},
		},
	}

	tests := []struct {
		Link     *Link
		URI      string
		Expected *ExpandedURL
	}{
		{
			&linkA,
			"a",
			&ExpandedURL{
				URL:   "https://a.com/",
				Index: 0,
				Link:  &linkA,
			},
		},
		{
			&linkA,
			"a/",
			&ExpandedURL{
				URL:   "https://a.com/",
				Index: 0,
				Link:  &linkA,
			},
		},
		{
			&linkA,
			"a/b/c/d/",
			&ExpandedURL{
				URL:   "https://a.com/b/c/d/",
				Index: 0,
				Link:  &linkA,
			},
		},
		{
			&linkB,
			"b/baz",
			nil,
		},
		{
			&linkB,
			"b/foo",
			&ExpandedURL{
				URL:   "https://b.com/foo",
				Index: 0,
				Link:  &linkB,
			},
		},
		{
			&linkB,
			"b/bar",
			&ExpandedURL{
				URL:   "https://b.com/bar",
				Index: 1,
				Link:  &linkB,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test-%d", i), func(t *testing.T) {
			if got := test.Link.Expand(test.URI); !expandedURLsAreSame(got, test.Expected) {
				t.Fatalf("got %v expected %v", got, test.Expected)
			}
		})
	}
}
