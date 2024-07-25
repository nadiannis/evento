package domain

type TicketTypeName string

var (
	TicketTypeVIP  TicketTypeName = "VIP"
	TicketTypeCAT1 TicketTypeName = "CAT 1"
)

type TicketType struct {
	ID    string         `json:"id"`
	Name  TicketTypeName `json:"name"`
	Price float64        `json:"price"`
}
