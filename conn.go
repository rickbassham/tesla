package tesla

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	// DefaultBaseURL is the URL for the Tesla owner's API.
	DefaultBaseURL = "https://owner-api.teslamotors.com"
)

// Conn represents a connection to the Tesla owner's API.
type Conn struct {
	rt           http.RoundTripper
	baseURL      string
	clientID     string
	clientSecret string

	accessToken  string
	refreshToken string

	debugMode bool
}

// NewConn creates a new connection.
func NewConn(rt http.RoundTripper, baseURL, clientID, clientSecret string) *Conn {
	return &Conn{
		rt:           rt,
		baseURL:      baseURL,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

// SetDebugMode turns debug mode on or off. If on, all requests and responses are dumped in their
// raw state to stdout.
func (c *Conn) SetDebugMode(debug bool) {
	if debug {
		c.rt = &debugTransport{
			r: c.rt,
		}
	} else if rt, ok := c.rt.(*debugTransport); ok {
		c.rt = rt.r
	}
}

// SetRefreshToken allows you to override the refresh token received from Authenticate.
func (c *Conn) SetRefreshToken(refreshToken string) {
	c.refreshToken = refreshToken
}

func (c *Conn) doRequest(method, url string, reqBody, respBody interface{}) error {
	var reqBodyReader io.Reader

	if reqBody != nil {
		reqBytes, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("error marshaling request body: %w", err)
		}

		reqBodyReader = bytes.NewReader(reqBytes)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.baseURL, url), reqBodyReader)
	if err != nil {
		return fmt.Errorf("error creating http request: %w", err)
	}

	if reqBodyReader != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	if c.accessToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	}

	resp, err := c.rt.RoundTrip(req)
	if err != nil {
		return fmt.Errorf("error performing http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w", HTTPStatusError{
			statusCode: resp.StatusCode,
		})
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading http response body: %w", err)
	}

	err = json.Unmarshal(respBytes, respBody)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}

	return nil
}
