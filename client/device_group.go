package client

import (
	models "gitlab.omnicube.ru/omnicube/omninanage/pkg/model/web"
	"omnilib/utils/converter"
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
	err = converter.SliceI2SliceModel(sources, &outSlice)
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
