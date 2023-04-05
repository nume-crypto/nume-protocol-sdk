package nume

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Auth struct {
	ApiKey string
}

const TIMEOUT = 10

type Request struct {
	Auth       Auth
	HTTPClient *http.Client
	Headers    map[string]string
	BaseURL    string
}

func (request *Request) doRequestResponse(req *http.Request) (map[string]interface{}, error) {
	client := request.HTTPClient
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resp := make(map[string]interface{})
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func buildURLWithParams(requestURL string, data map[string]interface{}) (string, error) {
	var URL *url.URL
	URL, err := url.Parse(requestURL)
	if err != nil {
		return "", err
	}
	parameters := url.Values{}
	for k, v := range data {
		parameters.Add(k, fmt.Sprintf("%v", v))
	}
	URL.RawQuery = parameters.Encode()
	return URL.String(), nil
}

func (request *Request) Post(path string, payload map[string]interface{}) (map[string]interface{}, error) {
	json, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s%s", request.BaseURL, path)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	req.Header.Set("NUME-API-KEY", request.Auth.ApiKey)
	return request.doRequestResponse(req)
}

func (request *Request) Get(path string, queryParams map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%s", request.BaseURL, path)
	url, err := buildURLWithParams(url, queryParams)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("NUME-API-KEY", request.Auth.ApiKey)
	return request.doRequestResponse(req)
}
