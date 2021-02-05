package client

import (
	"omnilib/models"
	"strconv"
)

type UserService struct {
	client *Client
}

func (s *UserService) GetList(companyId int) ([]*models.User, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/users/",
		new(models.User),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.User
	err = s.client.sourceSliceToOut(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *UserService) Get(id int) (*models.User, error) {
	model := new(models.User)
	if err := s.client.getSourceSingle(id, "/companies/@all/users/", model); err != nil {
		return nil, err
	}
	return model, nil
}
