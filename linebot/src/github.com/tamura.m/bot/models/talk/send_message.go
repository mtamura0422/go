package talk


import (
//	"bytes"
//	"encoding/json"
	"golang.org/x/net/context"
//	"io"
	"strings"
	"net/url"
//	"google.golang.org/appengine/log"
)


// ReplyMessageCall type
type ReplyMessageCall struct {
	c   *Client
	ctx context.Context
	apiKey string
	message string
}

// ReplyMessage method
func (client *Client) ReplyMessage(apiKey string, message string) *ReplyMessageCall {
	return &ReplyMessageCall{
		c:         client,
		apiKey:    apiKey,
		message:   message,
	}
}

// WithContext method
func (call *ReplyMessageCall) WithContext(ctx context.Context) *ReplyMessageCall {
	call.ctx = ctx
	return call
}





func (call *ReplyMessageCall) encodeParameter()(*strings.Reader) {

	/*
	enc := json.NewEncoder(w)

	return enc.Encode(&struct {
		Apikey string    `json:"apikey"`
		Query string    `json:"query"`
	}{
		Apikey: call.apiKey,
		Query:   call.message,
	})
*/

	// url.Valuesオブジェクト生成
	values := url.Values{}
	// key-valueを追加
  	values.Add("apikey", call.apiKey)
	values.Add("query", call.message) 

	return strings.NewReader(values.Encode())

}




// Do method
func (call *ReplyMessageCall) Do() (*TalkResponse, error) {

	body := call.encodeParameter()

	res, err := call.c.post(call.ctx, APIEndpointTalk, body)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeTalkResponse(res)
}