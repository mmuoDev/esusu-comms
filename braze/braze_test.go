package braze

import (
	http "net/http"
	"net/http/httptest"
	"testing"

	esusuEvent "github.com/mmuoDev/esusu-comms/events"
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
	// expected := `{
	// 	"message" : "success",
	// 	"attributes_processed" : 1
	// 	"events_processed" : 1
	// }`
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO: How do I serve an error here?
	}))
	defer svr.Close()
	c := NewProvider(apiKey, svr.URL)
	err := c.CustomEvent(testData())
	if err != nil {
		t.Errorf("expected res to be nil got %s", err)
	}
}
