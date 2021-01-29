package client

import (
	"omnilib/models"
	"strconv"
)

type UserService struct {
	client *Client
}

func (s *UserService) GetList(companyId int) ([]models.User, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/users/",
		&Source{Data: new(models.UserData), Relations: new(models.UserRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []models.User
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *UserService) Get(id int) (*models.User, error) {
	data := new(models.UserData)
	rel := new(models.UserRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/users/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &models.User{data, rel}, nil
}
