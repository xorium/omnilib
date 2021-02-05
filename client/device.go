package client

import (
	"omnilib/models"
	"strconv"
)

type DeviceService struct {
	client *Client
}

func (s *DeviceService) GetList(companyId int) ([]*models.Device, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/devices/",
		new(models.Device),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.Device
	err = s.client.sourceSliceToOut(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *DeviceService) Get(id int) (*models.Device, error) {
	model := new(models.Device)
	if err := s.client.getSourceSingle(id, "/companies/@all/devices/", model); err != nil {
		return nil, err
	}
	return model, nil
}
