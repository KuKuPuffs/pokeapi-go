package pokeapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const (
	// BaseURL ...
	BaseURL = "https://pokeapi.co/api/v2"
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
			Timeout: 1000 * time.Second,
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

	return &Client{
		HTTPClient: httpClient,
		BaseURL:    baseURL,
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
	rel, err := url.Parse(path)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing relative path URL")
	}

	pokeURL := fmt.Sprintf("%v%v", c.BaseURL.String(), rel.String())

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	fmt.Printf("URL to parse: %v\n", pokeURL)
	req, err := http.NewRequest(method, pokeURL, buf)
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

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("response code from server error, got: %d, want: %d", resp.StatusCode, http.StatusOK)
	}

	decodeErr := json.NewDecoder(resp.Body).Decode(v)
	return resp, decodeErr
}
