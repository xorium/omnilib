package client

import "omnilib/models"

type ParameterService struct {
	client *Client
}

func (s *ParameterService) GetList() ([]models.Parameter, error) {
	sources, err := s.client.getSourceMultiple(
		"/parameters/",
		&Source{Data: new(models.ParameterData)},
	)
	if err != nil {
		return nil, err
	}

	var out []models.Parameter
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *ParameterService) Get(id int) (*models.Parameter, error) {
	outData := new(models.ParameterData)
	if err := s.client.getSourceSingle(id, "/parameters/", &Source{Data: outData}); err != nil {
		return nil, err
	}
	return &models.Parameter{outData}, nil
}
