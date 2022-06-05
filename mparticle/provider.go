package mparticle

import (
	context "context"
	"errors"
	_nethttp "net/http"
	"strings"

	"github.com/mParticle/mparticle-go-sdk/events"
	esusuEvent "github.com/mmuoDev/esusu-comms/events"
)

const (
	DEV  = "DEV"
	PROD = "PROD"
)

type MParticleProvider interface {
	UploadEvents(ctx context.Context, batch events.Batch) (*_nethttp.Response, error)
}

type provider struct {
	apiKey            string
	secretKey         string
	env               string
	mparticleProvider MParticleProvider
}

func NewProvider(apiKey, secretKey, env string, mparticleProvider MParticleProvider) *provider {
	return &provider{
		apiKey:            apiKey,
		secretKey:         secretKey,
		env:               env,
		mparticleProvider: mparticleProvider,
	}
}

func (p *provider) CustomEvent(event esusuEvent.Event) error {
	batch := events.Batch{}
	ctx := context.WithValue(
		context.Background(),
		events.ContextBasicAuth,
		events.BasicAuth{
			APIKey:    p.apiKey,
			APISecret: p.secretKey,
		},
	)
	if strings.ToUpper(p.env) != DEV && strings.ToUpper(p.env) != PROD {
		return errors.New("A env var is required! - either DEV or PROD")
	}
	if strings.ToUpper(p.env) == PROD {
		batch.Environment = events.ProductionEnvironment
	} else {
		batch.Environment = events.DevelopmentEnvironment
	}
	//set user identities
	if _, ok := event.UserAttributes["external_id"]; ok {
		return errors.New("At least external_id is required for user identities")
	}
	email := ""
	val, exists := event.UserAttributes["email"]
	if exists {
		email = val.(string)
	}
	batch.UserIdentities = &events.UserIdentities{
		CustomerID: event.UserIdentities["external_id"],
		Email:      email,
	}
	//set user attributes
	batch.UserAttributes = make(map[string]interface{})
	for k, v := range event.UserAttributes {
		batch.UserAttributes[k] = v
	}
	//send custom events
	customEvent := events.NewCustomEvent()
	customEvent.Data.EventName = event.EventName
	customEvent.Data.CustomEventType = events.OtherCustomEventType
	customEvent.Data.CustomAttributes = make(map[string]interface{})
	for k, v := range event.CustomAttributes {
		customEvent.Data.CustomAttributes[k] = v
	}
	batch.Events = []events.Event{customEvent}
	_, err := p.mparticleProvider.UploadEvents(ctx, batch)
	if err != nil {
		return err
	}
	return nil
}
