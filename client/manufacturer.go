package client

import (
	models "github.com/xorium/omnimanage/pkg/model/web"
	"github.com/xorium/omnimanage/pkg/utils/converter"
)

type ManufacturerService struct {
	client *Client
}

func (s *ManufacturerService) GetList() ([]*models.Manufacturer, error) {

	sources, err := s.client.getSourceMultiple("/manufacturers/",
		new(models.Manufacturer),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.Manufacturer
	err = converter.SliceI2SliceModel(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *ManufacturerService) Get(id int) (*models.Manufacturer, error) {
	model := new(models.Manufacturer)
	if err := s.client.getSourceSingle(id, "/manufacturers/", model); err != nil {
		return nil, err
	}
	return model, nil
}
