package webhook

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	EventType      EventType `json:"eventType"`
	EventTime      time.Time `json:"eventTime"`
	OrganizationID uuid.UUID `json:"organizationId"`
	CorrelationID  uuid.UUID `json:"correlationId"`
}

type EventType string

const (
	EventTypeDeliveryOrderUpdate EventType = "DeliveryOrderUpdate"
	EventTypeDeliveryOrderError  EventType = "DeliveryOrderError"
	EventTypeReserveUpdate       EventType = "ReserveUpdate"
	EventTypeReserveError        EventType = "ReserveError"
	EventTypeTableOrderUpdate    EventType = "TableOrderUpdate"
	EventTypeTableOrderError     EventType = "TableOrderError"
	EventTypeStopListUpdate      EventType = "StopListUpdate"
	EventTypePersonalShiftUpdate EventType = "PersonalShift"
)
