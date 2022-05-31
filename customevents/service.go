package customevents

//CustomEventProvider represents method to be implemented in order to send custom events
type CustomEventProvider interface {
	CustomEvent (data interface{}) (interface{}, error)
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
func (s *service) Send(i interface{}) error {
	_, err := s.customEventProvider.CustomEvent(i)
	if err != nil {
		return err
	}
	//TODO: Check response
	return nil 
}