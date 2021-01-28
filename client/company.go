package client

type CompanyData struct {
	ID   int    `jsonapi:"primary,companies"`
	Name string `jsonapi:"attr,name"`
}

type Company struct {
	Data *CompanyData
}

type CompanyService struct {
	client *Client
}

func (s *CompanyService) GetList() ([]Company, error) {

	sources, err := s.client.getSourceMultiple("/companies/",
		&Source{Data: new(CompanyData)},
	)
	if err != nil {
		return nil, err
	}

	var outSlice []Company
	err = s.client.sourceSliceToOut(sources, &outSlice)

	return outSlice, nil
}

func (s *CompanyService) Get(id int) (*Company, error) {
	comp := new(CompanyData)
	if err := s.client.getSourceSingle(id, "/companies/", &Source{Data: comp}); err != nil {
		return nil, err
	}
	return &Company{Data: comp}, nil
}
