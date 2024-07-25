package domain

type Ticket struct {
	ID       string         `json:"id"`
	EventID  string         `json:"event_id"`
	Type     TicketTypeName `json:"type"`
	Quantity int            `json:"quantity"`
}
