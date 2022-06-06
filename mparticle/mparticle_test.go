package mparticle

import (
	"context"
	_nethttp "net/http"
	"testing"

	"github.com/mParticle/mparticle-go-sdk/events"
	esusuEvent "github.com/mmuoDev/esusu-comms/events"
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

func testDataWithoutExternalID() esusuEvent.Event {
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

type MockMparticle struct {
	uploadEvents func(ctx context.Context, batch events.Batch) (*_nethttp.Response, error)
}

func (m *MockMparticle) UploadEvents(ctx context.Context, batch events.Batch) (*_nethttp.Response, error) {
	if m.uploadEvents != nil {
		return m.uploadEvents(ctx, batch)
	}
	return &_nethttp.Response{}, nil
}

func TestMparticleWorksAsExpected(t *testing.T) {
	provider := NewProvider("", "", "dev", &MockMparticle{})
	err := provider.CustomEvent(testData())
	if err != nil {
		t.Fatal(err)
	}
}

