package restlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func interpretResponse(resp *http.Response, res interface{}) error {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, res)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	return err
}

// DoPost does the http POST call
func DoPost(url string, request interface{}, response interface{}) error {
	reqJson, err := json.Marshal(request)
	reader := bytes.NewReader(reqJson)
	resp, err := http.Post(url, "application/json", reader)
	if err != nil {
		panic(err)
	}
	return interpretResponse(resp, response)
}

// DoGet does the http GET call
func DoGet(url string, res interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return interpretResponse(resp, res)
}
