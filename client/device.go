package client

import (
	"omnilib/models"
	"strconv"
)

type DeviceService struct {
	client *Client
}

func (s *DeviceService) GetList(companyId int) ([]models.Device, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/devices/",
		&Source{Data: new(models.DeviceData), Relations: new(models.DeviceRelation)},
	)
	if err != nil {
		return nil, err
	}

	var devsOut []models.Device
	err = s.client.sourceSliceToOut(sources, &devsOut)

	return devsOut, nil
}

func (s *DeviceService) Get(id int) (*models.Device, error) {
	devData := new(models.DeviceData)
	devRel := new(models.DeviceRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/devices/", &Source{Data: devData, Relations: devRel}); err != nil {
		return nil, err
	}
	return &models.Device{Data: devData, Relations: devRel}, nil
}
