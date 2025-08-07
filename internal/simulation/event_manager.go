package simulation

import "fmt"

// EventType farklı olay türlerini belirtir
type EventType string

const (
	EventBirth    EventType = "birth"
	EventDeath    EventType = "death"
	EventConflict EventType = "conflict"
	EventNewYear  EventType = "new_year"
	// Diğer olay türleri buraya eklenebilir
)

// Event simülasyondaki bir olayı temsil eder
type Event struct {
	Type    EventType
	Message string
	Payload interface{} // Olayla ilgili ek veri (opsiyonel)
}

// EventManager olayları kaydeder ve işler
type EventManager struct {
	events []Event
}

// NewEventManager yeni bir EventManager oluşturur
func NewEventManager() *EventManager {
	return &EventManager{
		events: []Event{},
	}
}

// Emit yeni bir olay oluşturur ve listeye ekler
func (em *EventManager) Emit(eventType EventType, message string, payload interface{}) {
	event := Event{
		Type:    eventType,
		Message: message,
		Payload: payload,
	}
	em.events = append(em.events, event)
	fmt.Printf("[EVENT] %s: %s\n", eventType, message) // Konsola log
}

// GetEvents kayıtlı olayların listesini döner
func (em *EventManager) GetEvents() []Event {
	return em.events
}

// ClearEvents tüm olay listesini temizler
func (em *EventManager) ClearEvents() {
	em.events = []Event{}
}
