package tft

import "github.com/KnutZuidema/golio/internal"

// Client pools methods for the Legends of Runeterra API.
type Client struct {
	Match *MatchClient
}

// NewClient returns a new instance of a Legends of Runeterra client.
func NewClient(base *internal.Client) *Client {
	return &Client{
		Match: &MatchClient{c: base},
	}
}
