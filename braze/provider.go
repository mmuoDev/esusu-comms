package braze

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mmuoDev/esusu-comms/events"
)

const (
	endpoint = "https://sdk.iad-06.braze.com"
)

type event struct {
	name string
}

type provider struct {
	apiKey string
}

func NewProvider(apiKey string) *provider {
	return &provider{
		apiKey: apiKey,
	}
}

func (p *provider) CustomEvent(event events.Event) error {
	url := endpoint + "/users/track"
	client := &http.Client{}
	data := make(map[string]interface{})
	eventData := make(map[string]interface{})
	payload := make(map[string][]interface{})
	if len(event.UserIdentities) == 0 {
		return errors.New("At least external_id is required for user identities")
	}
	if event.EventName == "" {
		return errors.New("Event name is required!")
	}
	for k, v := range event.UserIdentities {
		data[k] = v
	}
	for k, v := range event.UserAttributes {
		data[k] = v
	}
	for k, v := range event.CustomAttributes {
		data[k] = v
	}
	eventData["name"] = event.EventName
	payload["attributes"] = []interface{}{data}
	payload["events"] = []interface{}{eventData}
	bb, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bb))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == 429 {
		delayString := res.Header.Get("Retry-After")
		delay := 30
		if delayString != "" {
			delay, err = strconv.Atoi(delayString)
			if err != nil {
				delay = 30
			}
		}
		time.Sleep(time.Duration(delay) * time.Second)
		return p.CustomEvent(event)
	}
	//TODO: Read body response
	return nil
}
