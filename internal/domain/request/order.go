package request

type OrderRequest struct {
	CustomerID string `json:"customer_id"`
	TicketID   string `json:"ticket_id"`
	Quantity   int    `json:"quantity"`
}
