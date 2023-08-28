package http

import (
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

type Response struct {
	raw *http.Response
}

func (r *Response) Json() (map[string]any, error) {
	resp, err := io.ReadAll(r.raw.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Response) String() (string, error) {
	resp, err := io.ReadAll(r.raw.Body)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (r *Response) StatusCode() int {
	return r.raw.StatusCode
}
