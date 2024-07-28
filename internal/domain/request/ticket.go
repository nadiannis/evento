package request

import "github.com/nadiannis/evento/internal/domain"

type TicketRequest struct {
	EventID  string                `json:"event_id"`
	Type     domain.TicketTypeName `json:"type"`
	Quantity int                   `json:"quantity"`
}

type TicketQuantityRequest struct {
	Quantity int `json:"quantity"`
}
