package restlib

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// HTTPResult is the result return by the library
type HTTPResult struct {
	HTTPResponse *http.Response
	Body         []byte
	Response     interface{}
}

func makeHTTPResult(res *http.Response, body []byte, resp interface{}) *HTTPResult {
	return &HTTPResult{
		HTTPResponse: res,
		Body:         body,
		Response:     resp,
	}
}

// DoHTTPRequest returns HTTPResult
func DoHTTPRequest(ctx context.Context, client *http.Client, method string, urlString string, body interface{}, headers map[string]string, required []string, response []interface{}) (*HTTPResult, error) {
	var reader io.Reader

	// Validations 1:
	// If we have body, marshal it to json
	if body != nil {
		reqJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(reqJSON)
	}

	// Validations 2:
	// if we have required headers, see if they have been passed to us
	for _, key := range required {
		if _, has := headers[key]; !has {
			return nil, errors.Errorf("Missing Required header: %s", key)
		}
	}

	httpRequest, err := http.NewRequest(method, urlString, reader)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		httpRequest.Header.Add(key, val)
	}

	httpResponse, err := client.Do(httpRequest.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	respBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	for _, respInterface := range response {
		err = json.Unmarshal(respBody, respInterface)
		if err == nil {
			return makeHTTPResult(httpResponse, respBody, respInterface), nil
		}
	}
	return makeHTTPResult(httpResponse, respBody, nil), nil
}
