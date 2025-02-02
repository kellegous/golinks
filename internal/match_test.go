package internal

import (
	"encoding/json"
	"errors"
	"regexp"
	"testing"
)

func errorsAreSame(a, b error) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Error() == b.Error()
}

func TestExpand(t *testing.T) {
	tests := []struct {
		Match       Match
		URI         string
		ExpectedURL string
		ExpectedOK  bool
	}{
		{
			// matching single capture unnamed
			Match{
				URIPattern:  regexp.MustCompile("^/a/(.*)$"),
				URLTemplate: "https://a.com/$1",
			},
			"/a/b/c",
			"https://a.com/b/c",
			true,
		},
		{
			// matching single capture named
			Match{
				URIPattern:  regexp.MustCompile("^/a/(?P<foo>.*)$"),
				URLTemplate: "https://a.com/$foo",
			},
			"/a/b/c",
			"https://a.com/b/c",
			true,
		},
		{
			// matching multiple captures unnamed
			Match{
				URIPattern:  regexp.MustCompile(`^/(\d+)/(\d+)$`),
				URLTemplate: "https://a.com/$2/$1",
			},
			"/12/13",
			"https://a.com/13/12",
			true,
		},
		{
			// matching multiple captures named
			Match{
				URIPattern:  regexp.MustCompile(`^/(?P<foo>\d+)/(?P<bar>\d+)$`),
				URLTemplate: "https://a.com/$bar/$foo",
			},
			"/12/13",
			"https://a.com/13/12",
			true,
		},
		{
			// unmatched
			Match{
				URIPattern:  regexp.MustCompile(`^/a(\d)$`),
				URLTemplate: "https://a.com/$1",
			},
			"/ab",
			"",
			false,
		},
		{
			// literal match
			Match{
				URIPattern:  regexp.MustCompile(`^literal$`),
				URLTemplate: "https://a.com",
			},
			"literal",
			"https://a.com",
			true,
		},
		{
			// literal match with unbounded captures
			Match{
				URIPattern:  regexp.MustCompile(`^literal$`),
				URLTemplate: "https://a.com/$1/$2",
			},
			"literal",
			"https://a.com//",
			true,
		},
	}

	for _, test := range tests {
		url, ok := test.Match.ExpandURL(test.URI)
		if url != test.ExpectedURL || ok != test.ExpectedOK {
			t.Fatalf(
				"match = %#v, uri = %#v, got (%#v, %t) expected (%#v, %t)",
				test.Match,
				test.URI,
				url,
				ok,
				test.ExpectedURL,
				test.ExpectedOK)
		}
	}
}

// TODO(kelly): Add tests for Match.UnmarshalJSON
func TestMatchUnmarshalJSON(t *testing.T) {
	tests := []struct {
		JSON          string
		ExpectedMatch *Match
		ExpectedError error
	}{
		{
			`{"uri_pattern":"^/a/(.*)$","url_template":"https://a.com/$1"}`,
			&Match{
				URIPattern:  regexp.MustCompile("^/a/(.*)$"),
				URLTemplate: "https://a.com/$1",
			},
			nil,
		},
		{
			`{"uri_pattern":"^/a/(?P<foo>.*)$","url_template":"https://a.com/$foo"}`,
			&Match{
				URIPattern:  regexp.MustCompile("^/a/(?P<foo>.*)$"),
				URLTemplate: "https://a.com/$foo",
			},
			nil,
		},
		{
			`{"uri_pattern":"^/a/(?P<foo>.*)$","url_template":"!://"}`,
			nil,
			errors.New("invalid URL"),
		},
		{
			`{"uri_pattern":"^/a/(?P<foo>.*)$","url_template":"mailto://a.com/$foo"}`,
			nil,
			errors.New("URL must be http or https"),
		},
		{
			`{"uri_pattern":"^/a/(?P<foo>.*)$","url_template":"https://$foo.com/"}`,
			nil,
			errors.New("URL host cannot contain '$' replacements"),
		},
	}

	for _, test := range tests {
		var m Match

		err := json.Unmarshal([]byte(test.JSON), &m)
		if !errorsAreSame(err, test.ExpectedError) {
			t.Fatalf("got error %v, expected %v", err, test.ExpectedError)
		}
		if test.ExpectedMatch != nil && !MatchesAreSame(&m, test.ExpectedMatch) {
			t.Fatalf("got %#v, expected %#v", m, test.ExpectedMatch)
		}
	}
}
