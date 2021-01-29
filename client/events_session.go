package client

import (
	"omnilib/models"
	"strconv"
)

type EventsSessionService struct {
	client *Client
}

func (s *EventsSessionService) GetList(companyId int) ([]models.EventsSession, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/events-sessions/",
		&Source{Data: new(models.EventsSessionData), Relations: new(models.EventsSessionRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []models.EventsSession
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *EventsSessionService) Get(id int) (*models.EventsSession, error) {
	data := new(models.EventsSessionData)
	rel := new(models.EventsSessionRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/events-sessions/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &models.EventsSession{data, rel}, nil
}
