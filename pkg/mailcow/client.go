package mailcow

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type MailcowClient struct {
	Host       string
	ApiKey     string
	httpClient http.Client
}

func NewClient(Host string, ApiKey string) MailcowClient {
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}
	return MailcowClient{Host: Host, ApiKey: ApiKey, httpClient: httpClient}
}

func (c MailcowClient) prepareRequest(endpoint string, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, c.Host+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-Key", c.ApiKey)
	return req, nil
}

func (c MailcowClient) GetMailbox(mailbox string) (*Mailbox, error) {
	req, err := c.prepareRequest("/api/v1/get/mailbox/"+mailbox, http.MethodGet)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	ret := Mailbox{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
