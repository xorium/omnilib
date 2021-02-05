package client

import (
	"gitlab.omnicube.ru/libs/omnilib/models"
	"strconv"
)

type DeviceGroupService struct {
	client *Client
}

func (s *DeviceGroupService) GetList(companyId int) ([]*models.DeviceGroup, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/device-groups/",
		new(models.DeviceGroup),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.DeviceGroup
	err = s.client.sourceSliceToOut(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *DeviceGroupService) Get(id int) (*models.DeviceGroup, error) {
	model := new(models.DeviceGroup)
	if err := s.client.getSourceSingle(id, "/companies/@all/device-groups/", model); err != nil {
		return nil, err
	}
	return model, nil
}
