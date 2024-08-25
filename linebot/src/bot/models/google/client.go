package google

import (
	"golang.org/x/net/context"
	"net/http"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/appengine/urlfetch"
)


// New returns a new bot client instance.
func New(c context.Context, api_key string) *http.Client {

	client := &http.Client{
		Transport: &transport.APIKey{
			Key:       api_key,
			Transport: &urlfetch.Transport{Context: c},
		},
    }

	return client
}

