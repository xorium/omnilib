package omnimlib

type ParameterData struct {
	ID                int                    `jsonapi:"primary,parameters"`
	Name              string                 `jsonapi:"attr,name"`
	Description       string                 `jsonapi:"attr,desc"`
	Type              string                 `jsonapi:"attr,type"`
	IsValuesSetFinite bool                   `jsonapi:"attr,isValuesSetFinite"`
	Info              map[string]interface{} `jsonapi:"attr,info"`
}

type Parameter struct {
	Data *ParameterData
}

type ParameterService struct {
	client *Client
}

func (s *ParameterService) GetList() ([]Parameter, error) {
	sources, err := s.client.getSourceMultiple(
		"/parameters/",
		&Source{Data: new(ParameterData)},
	)
	if err != nil {
		return nil, err
	}

	var out []Parameter
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *ParameterService) Get(id int) (*Parameter, error) {
	outData := new(ParameterData)
	if err := s.client.getSourceSingle(id, "/parameters/", &Source{Data: outData}); err != nil {
		return nil, err
	}
	return &Parameter{outData}, nil
}
