package nume

import (
	"net/http"
	"time"
)

type Client struct {
	User        *User
	Transaction *Transaction
}

var request *Request

func NewClient(api_key string) *Client {
	auth := Auth{ApiKey: api_key}
	httpClient := &http.Client{Timeout: TIMEOUT * time.Second}
	request = &Request{Auth: auth, HTTPClient: httpClient,
		BaseURL: BASE_URL}
	client := Client{
		User:        &User{Request: request},
		Transaction: &Transaction{Request: request},
	}
	return &client
}
