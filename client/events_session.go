package client

import (
	"omnilib/models"
	"strconv"
)

type EventsSessionService struct {
	client *Client
}

func (s *EventsSessionService) GetList(companyId int) ([]*models.EventsSession, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/events-sessions/",
		new(models.EventsSession),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.EventsSession
	err = s.client.sourceSliceToOut(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *EventsSessionService) Get(id int) (*models.EventsSession, error) {
	model := new(models.EventsSession)

	if err := s.client.getSourceSingle(id, "/companies/@all/events-sessions/", model); err != nil {
		return nil, err
	}
	return model, nil
}
