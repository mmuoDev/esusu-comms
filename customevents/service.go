package customevents

import (
	"github.com/mmuoDev/esusu-comms/events"
)

//CustomEventProvider represents method to be implemented in order to send custom events
type CustomEventProvider interface {
	CustomEvent(event events.Event) error
}

type service struct {
	customEventProvider CustomEventProvider
}

func NewService(c CustomEventProvider) *service {
	return &service{
		customEventProvider: c,
	}
}

//Send sends a custom event
func (s *service) CustomEvent(event events.Event) error {
	if err := s.customEventProvider.CustomEvent(event); err != nil {
		return err
	}
	return nil
}
