package helpscout

import (
	"net/url"
	"net/http"
	"strconv"
	"time"
	"io/ioutil"
	"encoding/json"
	"strings"
	"gitlab.com/diraven/wild-hub/src/github.com/go-errors/errors"
)

// Client manages communication with the HelpScout API.
type Client struct {
	httpClient *http.Client // HTTP httpClient used to communicate with the API.

	BaseURL *url.URL
	apiKey  string

	common service

	// Services used for talking to different parts of the GitHub API.
	Mailboxes *MailboxesService
	Conversations *ConversationsService
}

func New(apiKey string) *Client {
	baseUrl, _ := url.Parse("https://api.helpscout.net/v1")

	c := &Client{httpClient: http.DefaultClient, BaseURL: baseUrl}
	c.common.client = c
	c.apiKey = apiKey

	c.Mailboxes = (*MailboxesService)(&c.common)
	c.Conversations = (*ConversationsService)(&c.common)

	return c
}

func (c *Client) DoNewRequest(method, path string, result interface{}, params map[string]string) error {
	// Prepare necessary parameters.
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}

	// Generate request.
	req, err := c.NewRequest(http.MethodGet, path, values)
	if err != nil {
		return err
	}

	// Perform the request.
	err = c.Do(req, result)
	if err != nil {
		return err
	}

	// Return no errors on success.
	return nil
}

func (c *Client) NewRequest(method string, path string, params url.Values) (*http.Request, error) {
	// Get base url.
	theUrl := *c.BaseURL

	// Add necessary address parts into the base url.
	if strings.HasPrefix("/", path) {
		theUrl.Path = theUrl.Path + path
	} else {
		theUrl.Path = theUrl.Path + "/" + path
	}

	// Prepare request depending on it's type.
	var req *http.Request
	var err error
	if method == http.MethodGet {
		req, err = http.NewRequest(method, theUrl.String(), nil)
		if err != nil {
			return nil, err
		}
		req.URL.RawQuery = params.Encode()
	} else {
		req, err = http.NewRequest(method, theUrl.String(), strings.NewReader(params.Encode()))
		if err != nil {
			return nil, err
		}
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, result interface{}) error {
	// Authenticate request.
	req.SetBasicAuth(c.apiKey, "n/a")

	// Perform request.
	resp, err := c.httpClient.Do(req)

	// Process error response codes if any.
	if resp.StatusCode != 200 {
		return errors.New(strconv.Itoa(resp.StatusCode) + ": " + http.StatusText(resp.StatusCode))
	}

	// Check for api rate limiting.
	retryAfter := resp.Header.Get("Retry-After")
	retries := 0
	maxRetries := 10
	for retryAfter != "" && retries < maxRetries {
		// Retry request again after Retry-After seconds.
		waitSeconds, err := strconv.Atoi(retryAfter)
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(waitSeconds) * time.Second)
		resp, err = c.httpClient.Do(req)
		if err != nil {
			return err
		}
		retryAfter = resp.Header.Get("Retry-After")
		retries = retries + 1
	}

	// Read all the data from the response.
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal the response.
	err = json.Unmarshal(data, result)
	return err
}
