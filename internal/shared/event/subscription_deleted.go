package event

type SubscriptionDeleted struct {
	Name    string
	Payload interface{}
}

func NewSubscriptionDeleted() *SubscriptionDeleted {
	return &SubscriptionDeleted{
		Name: "SubscriptionDeleted",
	}
}

func (e SubscriptionDeleted) GetName() string {
	return e.Name
}

func (e SubscriptionDeleted) GetPayload() interface{} {
	return e.Payload
}

func (e *SubscriptionDeleted) SetPayload(payload interface{}) {
	e.Payload = payload
}
