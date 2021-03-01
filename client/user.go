package client

import (
	models "gitlab.omnicube.ru/omnicube/omninanage/pkg/model/web"
	"omnilib/utils/converter"
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
	err = converter.SliceI2SliceModel(sources, &outSlice)
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
