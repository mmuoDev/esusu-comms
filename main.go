package main

import (
	"log"

	"github.com/mmuoDev/esusu-comms/braze"
	"github.com/mmuoDev/esusu-comms/customevents"
	esusuEvent "github.com/mmuoDev/esusu-comms/events"
)

func main() {
	eventName := "TEST_CUSTOM_EVENT"
	ui := map[string]string{"external_id": "422c4d0b-7e7b-4033-95f1-3de455d1dc0b"}
	ua := map[string]interface{}{"site": "Holmes Homes"}
	ca := map[string]interface{}{"isCreditScoreRecorded": true}
	event := esusuEvent.Event{
		EventName:        eventName,
		UserIdentities:   ui,
		UserAttributes:   ua,
		CustomAttributes: ca,
	}
	brazeURL := "https://sdk.iad-06.braze.com"
	brazeSecretKey := ""

	brazeProvider := braze.NewProvider(brazeSecretKey, brazeURL)
	customEventService := customevents.NewService(brazeProvider)
	if err := customEventService.CustomEvent(event); err != nil {
		log.Fatal(err)
	}
}
