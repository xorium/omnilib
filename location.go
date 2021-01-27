package omnimlib

import "strconv"

type LocationData struct {
	ID       int                    `jsonapi:"primary,locations"`
	Name     string                 `jsonapi:"attr,name"`
	Timezone string                 `jsonapi:"attr,timezone"`
	Info     map[string]interface{} `jsonapi:"attr,info"`
}

type LocationRelation struct {
	Company  *CompanyData    `jsonapi:"relation,company"`
	Children []*LocationData `jsonapi:"relation,children"`
	Users    []*UserData     `jsonapi:"relation,users"`
}

type Location struct {
	Data      *LocationData
	Relations *LocationRelation
}

type LocationService struct {
	client *Client
}

func (s *LocationService) GetList(companyId int) ([]Location, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/locations/",
		&Source{Data: new(LocationData), Relations: new(LocationRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []Location
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *LocationService) Get(id int) (*Location, error) {
	data := new(LocationData)
	rel := new(LocationRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/locations/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &Location{data, rel}, nil
}
