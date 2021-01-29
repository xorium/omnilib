package client

import "omnilib/models"

type ManufacturerService struct {
	client *Client
}

func (s *ManufacturerService) GetList() ([]models.Manufacturer, error) {

	sources, err := s.client.getSourceMultiple("/manufacturers/",
		&Source{Data: new(models.ManufacturerData)},
	)
	if err != nil {
		return nil, err
	}

	var outSlice []models.Manufacturer
	err = s.client.sourceSliceToOut(sources, &outSlice)

	return outSlice, nil
}

func (s *ManufacturerService) Get(id int) (*models.Manufacturer, error) {
	data := new(models.ManufacturerData)
	if err := s.client.getSourceSingle(id, "/manufacturers/", &Source{Data: data}); err != nil {
		return nil, err
	}
	return &models.Manufacturer{Data: data}, nil
}
