package client

import (
	"strconv"
)

type DeviceGroupService struct {
	client *Client
}

func (s *DeviceGroupService) GetList(companyId int) ([]DeviceGroup, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/device-groups/",
		&Source{Data: new(DeviceGroupData), Relations: new(DeviceGroupRelation)},
	)
	if err != nil {
		return nil, err
	}

	var devsOut []DeviceGroup
	err = s.client.sourceSliceToOut(sources, &devsOut)

	return devsOut, nil
}

func (s *DeviceGroupService) Get(id int) (*DeviceGroup, error) {
	devData := new(DeviceGroupData)
	devRel := new(DeviceGroupRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/device-groups/", &Source{Data: devData, Relations: devRel}); err != nil {
		return nil, err
	}
	return &DeviceGroup{devData, devRel}, nil
}
