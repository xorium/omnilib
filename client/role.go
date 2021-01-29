package client

import (
	"omnilib/models"
	"strconv"
)

type RoleService struct {
	client *Client
}

func (s *RoleService) GetList(companyId int) ([]models.Role, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/roles/",
		&Source{Data: new(models.RoleData), Relations: new(models.RoleRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []models.Role
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *RoleService) Get(id int) (*models.Role, error) {
	data := new(models.RoleData)
	rel := new(models.RoleRelation)

	if err := s.client.getSourceSingle(id, "/companies/@all/roles/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &models.Role{data, rel}, nil
}
