package client

import "strconv"

type RoleData struct {
	ID         int                    `jsonapi:"primary,roles"`
	Name       string                 `jsonapi:"attr,name"`
	Slug       string                 `jsonapi:"attr,slug"`
	Assigned   bool                   `jsonapi:"attr,assigned"`
	Persistent bool                   `jsonapi:"attr,persistent"`
	Info       map[string]interface{} `jsonapi:"attr,info"`
}

type RoleRelation struct {
	Company *CompanyData `jsonapi:"relation,company"`
}

type Role struct {
	Data      *RoleData
	Relations *RoleRelation
}

type RoleService struct {
	client *Client
}

func (s *RoleService) GetList(companyId int) ([]Role, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/roles/",
		&Source{Data: new(RoleData), Relations: new(RoleRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []Role
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *RoleService) Get(id int) (*Role, error) {
	data := new(RoleData)
	rel := new(RoleRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/roles/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &Role{data, rel}, nil
}
