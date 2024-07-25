package main

import (
	"time"

	"github.com/nadiannis/evento/internal/domain/request"
	"github.com/nadiannis/evento/internal/usecase"
)

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

func prepopulateEvents(usecase usecase.IEventUsecase) {
	for _, eventInput := range eventInputs {
		usecase.Add(eventInput)
	}
}
