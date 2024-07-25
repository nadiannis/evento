package request

import "github.com/nadiannis/evento/internal/domain"

type TicketTypeRequest struct {
	Name  domain.TicketTypeName `json:"name"`
	Price float64               `json:"price"`
}
