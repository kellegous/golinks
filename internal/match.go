package internal

import (
	"encoding/json"
	"errors"
	"net/url"
	"regexp"
	"strings"
)

// Match represents a regexp pattern for matching the uri and, if matched, expands that match into a full URL using a given template.
type Match struct {
	URIPattern  *regexp.Regexp `json:"uri_pattern"`
	URLTemplate string         `json:"url_template"`
}

// ExpandURL expands the given uri into a full URL if the uri matches the pattern. If the uri does not match the pattern, the function returns false.
func (m *Match) ExpandURL(uri string) (string, bool) {
	p := m.URIPattern
	if idx := p.FindStringSubmatchIndex(uri); idx != nil {
		return string(p.ExpandString(nil, m.URLTemplate, uri, idx)), true
	}
	return "", false
}

func (m *Match) UnmarshalJSON(data []byte) error {
	var t struct {
		URIPattern  *regexp.Regexp `json:"uri_pattern"`
		URLTemplate string         `json:"url_template"`
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	if t.URIPattern == nil {
		return errors.New("uri_pattern missing")
	}

	if err := validateURL(t.URLTemplate); err != nil {
		return err
	}

	m.URIPattern = t.URIPattern
	m.URLTemplate = t.URLTemplate
	return nil
}

func validateURL(v string) error {
	u, err := url.Parse(v)
	if err != nil {
		return errors.New("invalid URL")
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("URL must be http or https")
	}

	if strings.Contains(u.Host, "$") {
		return errors.New("URL host cannot contain '$' replacements")
	}

	return nil
}

func MatchesAreSame(a, b *Match) bool {
	if a == b {
		return true
	} else if a == nil || b == nil {
		return false
	}

	return a.URIPattern.String() == b.URIPattern.String() &&
		a.URLTemplate == b.URLTemplate
}
