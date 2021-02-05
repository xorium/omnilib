package client

import "gitlab.omnicube.ru/libs/omnilib/models"

type ParameterService struct {
	client *Client
}

func (s *ParameterService) GetList() ([]*models.Parameter, error) {
	sources, err := s.client.getSourceMultiple(
		"/parameters/",
		new(models.Parameter),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.Parameter
	err = s.client.sourceSliceToOut(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *ParameterService) Get(id int) (*models.Parameter, error) {
	model := new(models.Parameter)
	if err := s.client.getSourceSingle(id, "/parameters/", model); err != nil {
		return nil, err
	}
	return model, nil
}
