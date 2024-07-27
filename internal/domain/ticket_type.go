package domain

type TicketTypeName string

var (
	TicketTypeVIP  TicketTypeName = "vip"
	TicketTypeCAT1 TicketTypeName = "cat1"
)

type TicketType struct {
	ID    string         `json:"id"`
	Name  TicketTypeName `json:"name"`
	Price float64        `json:"price"`
}
