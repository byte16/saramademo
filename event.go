package main

const (
	EVENT_TYPE_CREATE = "CreateEvent"
)

type CreateEvent struct {
	Type    string
	AccId   string
	AccName string
}

func (e CreateEvent) Process() (*BankAccount, error) {
	return updateAccount(e.AccId, map[string]interface{}{
		"Id":      e.AccId,
		"Name":    e.AccName,
		"Balance": "0",
	})
}

func NewCreateEvent(accName string) *CreateEvent {
	return &CreateEvent{
		Type:    EVENT_TYPE_CREATE,
		AccId:   UniqueId(),
		AccName: accName,
	}
}
