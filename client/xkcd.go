package client

import (
	"net/http"
	"time"
	"xkcd/model"
)

type ComicNumber int

const (
	// BaseURL for xkcd comics
	BaseURL string = "https://xkcd.com"
	// DefaultClientTimeout for
	DefaultClientTimeout time.Duration = 30 * time.Second
	//latest comic is the most recent comic per the API
	LatestComic ComicNumber = 0
)

type XKCDClient struct {
	client  *http.Client
	baseURL string
}

// XKCDClient is the client for XKCD
func NewXKCDClient() *XKCDClient {
	return &XKCDClient{
		client: &http.Client{
			Timeout: DefaultTimeout,
		},
		baseURL: BaseURL,
	}
}

// SetTimeout overrides the default ClientTimeout
func (hc *XKCDClient) SetTimeout(d time.Duration) {
	hc.client.Timeout = d
}

// Fetch retrieves the comic as per provided comic number
func (hc *XKCDClient) Fetch(n ComicNumber, save bool) (model.Comic, error) {
	resp, err := hc.client.Get(hc.buildURL(n))
	if err != nil {
		return model.Comic{}, err
	}
}
