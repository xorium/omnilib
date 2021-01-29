package client

import "omnilib/models"

type DeviceModelService struct {
	client *Client
}

func (s *DeviceModelService) GetList() ([]models.DeviceModel, error) {
	sources, err := s.client.getSourceMultiple(
		"/device-models/",
		&Source{Data: new(models.DeviceModelData), Relations: new(models.DeviceModelRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []models.DeviceModel
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *DeviceModelService) Get(id int) (*models.DeviceModel, error) {
	data := new(models.DeviceModelData)
	rel := new(models.DeviceModelRelation)
	if err := s.client.getSourceSingle(id, "/device-models/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &models.DeviceModel{data, rel}, nil
}
