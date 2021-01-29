package client

import (
	"omnilib/models"
	"strconv"
)

type Event struct {
	Data      *models.EventData
	Relations *models.EventRelation
}

type EventService struct {
	client *Client
}

func (s *EventService) GetList(companyId int) ([]Event, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/events/",
		&Source{Data: new(models.EventData), Relations: new(models.EventRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []Event
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *EventService) Get(id int) (*Event, error) {
	data := new(models.EventData)
	rel := new(models.EventRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/events/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &Event{data, rel}, nil
}
