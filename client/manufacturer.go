package client

type ManufacturerData struct {
	ID   int                    `jsonapi:"primary,manufacturers"`
	Name string                 `jsonapi:"attr,name"`
	Info map[string]interface{} `jsonapi:"attr,info"`
}

type Manufacturer struct {
	Data *ManufacturerData
}

type ManufacturerService struct {
	client *Client
}

func (s *ManufacturerService) GetList() ([]Manufacturer, error) {

	sources, err := s.client.getSourceMultiple("/manufacturers/",
		&Source{Data: new(ManufacturerData)},
	)
	if err != nil {
		return nil, err
	}

	var outSlice []Manufacturer
	err = s.client.sourceSliceToOut(sources, &outSlice)

	return outSlice, nil
}

func (s *ManufacturerService) Get(id int) (*Manufacturer, error) {
	data := new(ManufacturerData)
	if err := s.client.getSourceSingle(id, "/manufacturers/", &Source{Data: data}); err != nil {
		return nil, err
	}
	return &Manufacturer{Data: data}, nil
}
