package buienradar

import (
	"net/http"
	"encoding/json"
)

// Client is a HTTP client for the Buienradar API
type Client struct {
	*http.Client
}

// NewClient creates a new HTTP client for the Buienradar API
func NewClient(cli *http.Client) *Client {
	return &Client{	Client: cli}
}

// Get retrieves the latest data from the Buienradar API
func (c *Client) Get() (b *Buienradar, err error) {
	b = new(Buienradar)

	var resp *http.Response
	resp, err = c.Client.Get(URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(b)
	return
}

// Get retrieves the latest data from the Buienradar API
func Get() (b *Buienradar, err error) {
	return NewClient(&http.Client{}).Get()
}
