package client

import (
	"omnilib/models"
	"strconv"
)

type DeviceGroupService struct {
	client *Client
}

func (s *DeviceGroupService) GetList(companyId int) ([]models.DeviceGroup, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/device-groups/",
		&Source{Data: new(models.DeviceGroupData), Relations: new(models.DeviceGroupRelation)},
	)
	if err != nil {
		return nil, err
	}

	var devsOut []models.DeviceGroup
	err = s.client.sourceSliceToOut(sources, &devsOut)

	return devsOut, nil
}

func (s *DeviceGroupService) Get(id int) (*models.DeviceGroup, error) {
	devData := new(models.DeviceGroupData)
	devRel := new(models.DeviceGroupRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/device-groups/", &Source{Data: devData, Relations: devRel}); err != nil {
		return nil, err
	}
	return &models.DeviceGroup{devData, devRel}, nil
}
