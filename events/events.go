package events

//userIdentities is a type for user identities
type userIdentities map[string]string

//userAttributes is a type for user attributes
type userAttributes map[string]interface{}

//customAttributes is a type for custom attributes
type customAttributes map[string]interface{}

type Event struct {
	EventName        string
	UserIdentities   userIdentities
	UserAttributes   userAttributes
	CustomAttributes customAttributes
}
