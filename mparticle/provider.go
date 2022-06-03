package mparticle

import (
	context "context"
	_nethttp "net/http"

	"github.com/mParticle/mparticle-go-sdk/events"
)

type MParticleProvider interface {
	UploadEvents(ctx context.Context, batch events.Batch) (*_nethttp.Response, error) 
}

type provider struct {
	apiKey string
	secretKey string
	mparticleProvider MParticleProvider
}

func NewProvider(apiKey, secretKey string, mparticleProvider MParticleProvider) *provider {
	return &provider{
		apiKey: apiKey,
		secretKey: secretKey,
		mparticleProvider: mparticleProvider,
	}
}

func (p *provider) CustomEvent(data interface{}) error {
	ctx := context.WithValue(
		context.Background(),
		events.ContextBasicAuth,
		events.BasicAuth{
			APIKey:    p.apiKey,
			APISecret: p.secretKey,
		},
	)
	res, err := p.mparticleProvider.UploadEvents(ctx, data.(events.Batch))
	if err != nil {
		return err
	}
	if res != nil && res.StatusCode == 202 {
		return nil 
	}
	return err
}


