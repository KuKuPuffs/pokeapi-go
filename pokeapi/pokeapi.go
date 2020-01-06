package pokeapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	// BaseURL ...
	BaseURL = "https://pokeapi.co/api/v2/"
	// UserAgent ...
	UserAgent = "poke-go"
)

// Client struct is to initialize your HTTP Client.
type Client struct {
	HTTPClient *http.Client
	BaseURL    *url.URL
	UserAgent  string
}

// NewClient ...
func NewClient(httpClient *http.Client) *Client {
	if httpClient.Transport == nil {
		httpClient = &http.Client{
			Timeout: 100 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}
	baseURL, err := url.Parse(BaseURL)
	if err != nil {
		log.Fatal("can't parse Poke URL")
	}
	// fmt.Printf("baseURL: %v\n", BaseURL)
	pokeURL := baseURL.ResolveReference(baseURL)
	// fmt.Printf("pokeURL: %v\n", pokeURL)

	return &Client{
		HTTPClient: httpClient,
		BaseURL:    pokeURL,
		UserAgent:  UserAgent,
	}
}

func NewClientWIthOpts(opts ...func(*Client)) *Client {
	httpClient := &Client{
		HTTPClient: &http.Client{Timeout: time.Millisecond * 100000,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			}}},
		UserAgent: UserAgent,
	}

	for _, opt := range opts {
		opt(httpClient)
	}

	return httpClient
}

// OptionBaseURL sets the url for the client. only useful for testing.
func OptionBaseURL(u string) func(*Client) {
	return func(c *Client) {
		baseURL, err := url.Parse(u)
		if err != nil {
			log.Printf("url.Parse(%v) is unable to parse requested URL, err: %v", u, err)
		}
		c.BaseURL = baseURL
	}
}

// OptionHTTPClient sets the http.Client for the client field. only useful for testing.
func OptionHTTPClient(client *http.Client) func(*Client) {
	return func(c *Client) {
		c.HTTPClient = client
	}
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	relPath := &url.URL{Path: path}
	URL := c.BaseURL.ResolveReference(relPath)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
	}()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
