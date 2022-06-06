package braze

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/mmuoDev/esusu-comms/events"
)

type event struct {
	name string
}

type provider struct {
	apiKey string
	url    string
}

func NewProvider(apiKey, url string) *provider {
	return &provider{
		apiKey: apiKey,
		url:    url,
	}
}

func (p *provider) MakeHTTPCall(event events.Event) (*http.Response, error) {
	url := p.url + "/users/track"
	//url := p.url
	client := &http.Client{}
	data := make(map[string]interface{})
	eventData := make(map[string]interface{})
	payload := make(map[string][]interface{})
	if _, ok := event.UserIdentities["external_id"]; !ok {
		return nil, errors.New("At least external_id is required for user identities")
	}
	if event.EventName == "" {
		return nil, errors.New("Event name is required!")
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
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bb))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	//res, _ := client.Do(req)
	// bb, er := ioutil.ReadAll(res.Body)
	// if er != nil {
	// 	log.Fatal(er)
	// }
	// bodyString := string(bb)
	// log.Fatal(bodyString, res.StatusCode)
	return client.Do(req)
}

func (p *provider) CustomEvent(event events.Event) error {
	//url := p.url + "/users/track"
	// client := &http.Client{}
	// data := make(map[string]interface{})
	// eventData := make(map[string]interface{})
	// payload := make(map[string][]interface{})
	// if _, ok := event.UserIdentities["external_id"]; !ok {
	// 	return errors.New("At least external_id is required for user identities")
	// }
	// if event.EventName == "" {
	// 	return errors.New("Event name is required!")
	// }
	// for k, v := range event.UserIdentities {
	// 	data[k] = v
	// }
	// for k, v := range event.UserAttributes {
	// 	data[k] = v
	// }
	// for k, v := range event.CustomAttributes {
	// 	data[k] = v
	// }
	// eventData["name"] = event.EventName
	// payload["attributes"] = []interface{}{data}
	// payload["events"] = []interface{}{eventData}
	// bb, err := json.Marshal(payload)
	// if err != nil {
	// 	return err
	// }
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(bb))
	// if err != nil {
	// 	return err
	// }
	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	// res, err := client.Do(req)
	// if err != nil {
	// 	return err
	// }
	res, err := p.MakeHTTPCall(event)
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
	if res.StatusCode != 201 {
		bb, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("response didn't return a 201 status, unable to read braze's body response due to=%v", err)
		}
		bodyString := string(bb)
		return fmt.Errorf("response didn't return a 201 status due to=%v", bodyString)
	}
	return nil
}
