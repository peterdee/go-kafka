package events

type clientEvent struct {
	EventTarget string `json:"eventTarget"`
	EventType   string `json:"eventType"`
}
