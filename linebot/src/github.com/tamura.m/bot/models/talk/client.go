package talk

import (
	"io"
	"net/url"
	"net/http"
	"path"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
 //	"google.golang.org/appengine/log"

)


type TalkReply struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Results []struct {
		Perplexity float64 `json:"perplexity"`
		Reply string `json:"reply"`
	} `json:"results"`
}


// Client type
type Client struct {
	endpointBase  *url.URL     // default APIEndpointBase
	httpClient    *http.Client // default http.DefaultClient
}

// ClientOption type
type ClientOption func(*Client) error

// APIEndpoint constants
const (
	APIEndpointBase = "https://api.a3rt.recruit-tech.co.jp"
	APIEndpointTalk = "/talk/v1/smalltalk"
)



// New returns a new bot client instance.
func New(options ...ClientOption) (*Client, error) {
	c := &Client{
		httpClient: http.DefaultClient,
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	if c.endpointBase == nil {
		u, err := url.ParseRequestURI(APIEndpointBase)
		if err != nil {
			return nil, err
		}
		c.endpointBase = u
	}
	return c, nil
}



// WithHTTPClient function
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c
		return nil
	}
}



func (client *Client) url(endpoint string) string {
	u := *client.endpointBase
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}


func (client *Client) post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest("POST", client.url(endpoint), body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	return client.do(ctx, req)
}



func (client *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx != nil {
		return ctxhttp.Do(ctx, client.httpClient, req)
	}
	return client.httpClient.Do(req)

}



