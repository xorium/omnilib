package omnimlib

import (
	"strconv"
)

type DeviceGroupData struct {
	ID          int                    `jsonapi:"primary,deviceGroups"`
	Name        string                 `jsonapi:"attr,name"`
	Description string                 `jsonapi:"attr,desc"`
	Type        string                 `jsonapi:"attr,type"`
	Filters     map[string]interface{} `jsonapi:"attr,filters"`
}

type DeviceGroupRelation struct {
	Company *CompanyData  `jsonapi:"relation,company"`
	Devices []*DeviceData `jsonapi:"relation,devices"`
	User    *UserData     `jsonapi:"relation,user"`
}

type DeviceGroup struct {
	Data      *DeviceGroupData
	Relations *DeviceGroupRelation
}

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
