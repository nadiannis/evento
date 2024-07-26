package main

import (
	"time"

	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/domain/request"
	"github.com/nadiannis/evento/internal/usecase"
)

var ticketTypeInputs = []*request.TicketTypeRequest{
	{
		Name:  domain.TicketTypeVIP,
		Price: 5000,
	},
	{
		Name:  domain.TicketTypeVIP,
		Price: 250,
	},
}

var eventInputs = []*request.EventRequest{
	{
		Name: "Event 1",
		Date: time.Now().AddDate(0, 0, 14),
	},
	{
		Name: "Event 2",
		Date: time.Now().AddDate(0, 1, 0),
	},
	{
		Name: "Event 3",
		Date: time.Now().AddDate(0, 1, 14),
	},
	{
		Name: "Event 4",
		Date: time.Now().AddDate(0, 2, 0),
	},
	{
		Name: "Event 5",
		Date: time.Now().AddDate(0, 2, 14),
	},
}

func prepopulateTicketTypes(usecase usecase.ITicketTypeUsecase) {
	for _, ticketTypeInput := range ticketTypeInputs {
		usecase.Add(ticketTypeInput)
	}
}

func prepopulateEventsAndTickets(eventUsecase usecase.IEventUsecase, ticketUsecase usecase.ITicketUsecase) {
	for _, eventInput := range eventInputs {
		event := eventUsecase.Add(eventInput)

		vipTicket := &request.TicketRequest{
			EventID:  event.ID,
			Type:     domain.TicketTypeVIP,
			Quantity: 10,
		}
		cat1Ticket := &request.TicketRequest{
			EventID:  event.ID,
			Type:     domain.TicketTypeCAT1,
			Quantity: 100,
		}
		ticketUsecase.Add(vipTicket)
		ticketUsecase.Add(cat1Ticket)
	}
}
