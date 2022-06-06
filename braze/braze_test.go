package braze

import (
	"io/ioutil"
	http "net/http"
	"net/http/httptest"
	"testing"

	esusuEvent "github.com/mmuoDev/esusu-comms/events"
	"github.com/stretchr/testify/assert"
)

const (
	apiKey = "73585"
)

func testData() esusuEvent.Event {
	eventName := "TEST_CUSTOM_EVENT"
	ui := map[string]string{"external_id": "311c4d0b-7e7b-4033-95f1-3de455d1dc0b"}
	ua := map[string]interface{}{"site": "Holmes Homes"}
	ca := map[string]interface{}{"isCreditScoreRecorded": true}
	event := esusuEvent.Event{
		EventName:        eventName,
		UserIdentities:   ui,
		UserAttributes:   ua,
		CustomAttributes: ca,
	}
	return event
}
func TestBrazeCall(t *testing.T) {
	expected := `
	{
		"message" : "success",
		"attributes_processed" : 1
		"events_processed" 1
	  }
	`
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(expected))
		w.WriteHeader(http.StatusCreated)
	}))
	defer svr.Close()
	c := NewProvider(apiKey, svr.URL)
	res, err := c.MakeHTTPCall(testData())
	if err != nil {
		t.Errorf("expected res to be nil got %s", err)
	}
	bb, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	bs := string(bb)
	t.Run("Body is as expected", func(t *testing.T) {
		assert.Equal(t, expected, bs)
	})
	t.Run("HTTP status is as expected", func(t *testing.T) {
		assert.Equal(t, http.StatusCreated, 201)
	})
}
