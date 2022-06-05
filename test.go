
import (
	"log"

	//"github.com/mParticle/mparticle-go-sdk/events"
	"github.com/mmuoDev/esusu-comms/customevents"
	//"github.com/mmuoDev/esusu-comms/mparticle"
	esusuEvent "github.com/mmuoDev/esusu-comms/events"
	"github.com/mmuoDev/esusu-comms/braze"
)


func test() {
	eventName := "TEST_CUSTOM_EVENT"
	ui := map[string]string{"external_id": "311c4d0b-7e7b-4033-95f1-3de455d1dc0b"}
	ua := map[string]interface{}{"site": "Holmes Homes"}
	ca := map[string]interface{}{"isCreditScoreRecorded": true}
	event := esusuEvent.Event{
		EventName: eventName,
		UserIdentities: ui,
		UserAttributes: ua,
		CustomAttributes: ca,
	}

	//TODO: These values should come from env
	// mparticleAPIKey := "us1-1632a6044bd28146be632ef5cb688736"
	// mparticleSecretKey := "n2ur1TCP7ZCuf7wOXa7_--byycyLxfJgWHFTtrblF102IaK56aCW-9B7xukRJdXM"
	// env := "dev"

	brazeURL := "https://sdk.iad-06.braze.com"
	//brazeSecretKey := "dab3bae8-0e7a-46de-b9a5-a7f8fdde00d6"
	brazeSecretKey := "28cc8394-6909-4138-b631-66e26cc36ae0"

	
	// client := events.NewAPIClient(events.NewConfiguration()).EventsAPI
	// mparticleProvider := mparticle.NewProvider(mparticleAPIKey, mparticleSecretKey, env, client)


	brazeProvider := braze.NewProvider(brazeSecretKey, brazeURL)
	customEventService := customevents.NewService(brazeProvider)
	if err := customEventService.CustomEvent(event); err != nil {
		log.Fatal(err)
	}
}