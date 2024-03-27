package event

type SubscriptionSaved struct {
	Name    string
	Payload interface{}
}

func NewSubscriptionSaved() *SubscriptionSaved {
	return &SubscriptionSaved{
		Name: "SubscriptionSaved",
	}
}

func (e SubscriptionSaved) GetName() string {
	return e.Name
}

func (e SubscriptionSaved) GetPayload() interface{} {
	return e.Payload
}

func (e *SubscriptionSaved) SetPayload(payload interface{}) {
	e.Payload = payload
}
