package internal

type ExpandedURL struct {
	URL   string `json:"url"`
	Index int    `json:"index"`
	Link  *Link  `json:"match,omitempty"`
}
