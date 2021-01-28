package omnimlib

type DeviceModelData struct {
	ID          int    `jsonapi:"primary,deviceModels"`
	Name        string `jsonapi:"attr,name"`
	Title       string `jsonapi:"attr,title"`
	Description string `jsonapi:"attr,desc"`
}

type DeviceModelRelation struct {
	Manufacturer *ManufacturerData `jsonapi:"relation,manufacturer"`
}

type DeviceModel struct {
	Data      *DeviceModelData
	Relations *DeviceModelRelation
}

type DeviceModelService struct {
	client *Client
}

func (s *DeviceModelService) GetList() ([]DeviceModel, error) {
	sources, err := s.client.getSourceMultiple(
		"/device-models/",
		&Source{Data: new(DeviceModelData), Relations: new(DeviceModelRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []DeviceModel
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *DeviceModelService) Get(id int) (*DeviceModel, error) {
	data := new(DeviceModelData)
	rel := new(DeviceModelRelation)
	if err := s.client.getSourceSingle(id, "/device-models/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &DeviceModel{data, rel}, nil
}
