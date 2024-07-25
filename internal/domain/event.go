package domain

import "time"

type Event struct {
	ID      string                     `json:"id"`
	Name    string                     `json:"name"`
	Date    time.Time                  `json:"date"`
	Tickets map[TicketTypeName]*Ticket `json:"tickets"`
}
