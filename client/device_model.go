package client

import "gitlab.omnicube.ru/libs/omnilib/models"

type DeviceModelService struct {
	client *Client
}

func (s *DeviceModelService) GetList() ([]*models.DeviceModel, error) {
	sources, err := s.client.getSourceMultiple(
		"/device-models/",
		new(models.DeviceModel),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.DeviceModel
	err = s.client.sourceSliceToOut(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *DeviceModelService) Get(id int) (*models.DeviceModel, error) {
	model := new(models.DeviceModel)
	if err := s.client.getSourceSingle(id, "/device-models/", model); err != nil {
		return nil, err
	}
	return model, nil
}
