package client

import (
	"gitlab.omnicube.ru/libs/omnilib/models"
	"strconv"
)

type RoleService struct {
	client *Client
}

func (s *RoleService) GetList(companyId int) ([]*models.Role, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/roles/",
		new(models.Role),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.Role
	err = s.client.sourceSliceToOut(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *RoleService) Get(id int) (*models.Role, error) {
	model := new(models.Role)
	if err := s.client.getSourceSingle(id, "/companies/@all/roles/", model); err != nil {
		return nil, err
	}
	return model, nil
}
