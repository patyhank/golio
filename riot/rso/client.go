package rso

import "github.com/KnutZuidema/golio/internal"

type Client struct {
	Account *AccountClient
}

// NewClient returns a new instance of a League of Legends client.
func NewClient(base *internal.Client) *Client {
	return &Client{
		Account: &AccountClient{c: base},
	}
}
