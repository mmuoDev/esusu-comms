package braze

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	endpoint = "https://sdk.iad-06.braze.com"
)

type provider struct {
	apiKey string
}

func NewProvider(apiKey string) *provider {
	return &provider{
		apiKey: apiKey,
	}
}

func (p *provider) CustomEvent(data interface{}) (interface{}, error) {
	url := endpoint + "/users/track"
	client := &http.Client{}
	payload := strings.NewReader(data.(string))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
