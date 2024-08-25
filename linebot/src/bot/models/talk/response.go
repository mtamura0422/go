package talk

import (
	"encoding/json"
	"net/http"
)

// BasicResponse type
type BasicResponse struct {
}

type errorResponseDetail struct {
	Message  string `json:"message"`
	Property string `json:"property"`
}

// ErrorResponse type
type ErrorResponse struct {
	Message string                `json:"message"`
	Details []errorResponseDetail `json:"details"`
}



type TalkResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Results []struct {
		Perplexity float64 `json:"perplexity"`
		Reply string `json:"reply"`
	} `json:"results"`
}


func checkResponse(res *http.Response) error {
	if res.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(res.Body)
		result := ErrorResponse{}
		if err := decoder.Decode(&result); err != nil {
			return &APIError{
				Code: res.StatusCode,
			}
		}
		return &APIError{
			Code:     res.StatusCode,
			Response: &result,
		}
	}
	return nil
}

func decodeToBasicResponse(res *http.Response) (*BasicResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := BasicResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeTalkResponse(res *http.Response) (*TalkResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)

	result := TalkResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

