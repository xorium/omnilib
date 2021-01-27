package omnimlib

import (
	"strconv"
)

type DeviceData struct {
	ID          int                    `jsonapi:"primary,devices"`
	Name        string                 `jsonapi:"attr,name"`
	Slug        string                 `jsonapi:"attr,slug"`
	Title       string                 `jsonapi:"attr,title"`
	Description string                 `jsonapi:"attr,desc"`
	Kind        string                 `jsonapi:"attr,kind"`
	Info        map[string]interface{} `jsonapi:"attr,info"`
	Image       string                 `jsonapi:"attr,image"`

	//Info        struct {
	//	ToolsNum     int `jsonapi:"attr,toolsNum"`
	//	OperatorsNum int `jsonapi:"attr,operatorsNum"`
	//	Connection   struct {
	//		IP                 string `jsonapi:"attr,ip"`
	//		Port               int    `jsonapi:"attr,port"`
	//		ConnectingInterval int    `jsonapi:"attr,connecting_interval"`
	//		PingTimeoutSec     int    `jsonapi:"attr,ping_timeout_sec"`
	//		RecvLengthByte     int    `jsonapi:"attr,recv_length_byte"`
	//		RecvTerminator     string `jsonapi:"attr,recv_terminator"`
	//	} `jsonapi:"attr,connection"`
	//} `jsonapi:"attr,info"`
}

type DeviceRelation struct {
	Company  *CompanyData       `jsonapi:"relation,company"`
	Model    *DeviceModelData   `jsonapi:"relation,model"`
	Location *LocationData      `jsonapi:"relation,location"`
	Groups   []*DeviceGroupData `jsonapi:"relation,groups"`
	Parent   *DeviceData        `jsonapi:"relation,parent"`
	Rules    []*RuleData        `jsonapi:"relation,rules"`
}

type Device struct {
	Data      *DeviceData
	Relations *DeviceRelation
}

type DeviceService struct {
	client *Client
}

func (s *DeviceService) GetList(companyId int) ([]Device, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/devices/",
		&Source{Data: new(DeviceData), Relations: new(DeviceRelation)},
	)
	if err != nil {
		return nil, err
	}

	var devsOut []Device
	err = s.client.sourceSliceToOut(sources, &devsOut)

	return devsOut, nil
}

func (s *DeviceService) Get(id int) (*Device, error) {
	devData := new(DeviceData)
	devRel := new(DeviceRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/devices/", &Source{Data: devData, Relations: devRel}); err != nil {
		return nil, err
	}
	return &Device{devData, devRel}, nil
}
