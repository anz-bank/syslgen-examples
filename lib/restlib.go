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

// HttpResult is the result return by the library
type HttpResult struct {
	HttpResponse *http.Response
	Body         []byte
	Response     interface{}
}

func makeHttpResult(res *http.Response, body []byte, resp interface{}) *HttpResult {
	return &HttpResult{
		HttpResponse: res,
		Body:         body,
		Response:     resp,
	}
}

// DoHttpRequest returns HttpResult
// with HttpResult.httpResponse.Body set to nil
func DoHttpRequest(client *http.Client, ctx context.Context, method string, urlString string, body interface{}, headers map[string]string, required []string, response []interface{}) (*HttpResult, error) {
	var reader io.Reader
	var err error

	// Validations 1:
	// If we have body, marshal it to json
	if body != nil {
		reqJson, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(reqJson)
	}

	// Validations 2:
	// if we have required headers, see if they have been passed to us
	for _, key := range required {
		if _, has := headers[key]; !has {
			return nil, errors.Errorf("Missing Required header: %s", key)
		}
	}

	httpRequest, err := http.NewRequest(method, urlString, reader)

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
			return makeHttpResult(httpResponse, respBody, respInterface), nil
		}
	}
	return makeHttpResult(httpResponse, respBody, nil), nil
}
