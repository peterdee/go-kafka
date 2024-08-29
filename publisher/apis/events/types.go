package events

type ClientEvent struct {
	EventTarget string `json:"eventTarget"`
	EventType   string `json:"eventType"`
}
