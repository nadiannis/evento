package domain

import "time"

type Order struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	TicketID   string    `json:"ticket_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}
