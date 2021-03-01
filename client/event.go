package client

import (
	models "gitlab.omnicube.ru/omnicube/omninanage/pkg/model/web"
	"omnilib/utils/converter"
	"strconv"
)

type EventService struct {
	client *Client
}

func (s *EventService) GetList(companyId int) ([]*models.Event, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/events/",
		new(models.Event),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.Event
	err = converter.SliceI2SliceModel(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *EventService) Get(id int) (*models.Event, error) {
	model := new(models.Event)
	if err := s.client.getSourceSingle(id, "/companies/@all/events/", model); err != nil {
		return nil, err
	}
	return model, nil
}
