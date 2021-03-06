package client

import (
	models "github.com/xorium/omnimanage/pkg/model/web"
	"github.com/xorium/omnimanage/pkg/utils/converter"
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
	err = converter.SliceI2SliceModel(sources, &outSlice)
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
