package client

import (
	models "gitlab.omnicube.ru/omnicube/omninanage/pkg/model/web"
	"omnilib/utils/converter"
	"strconv"
)

type LocationService struct {
	client *Client
}

func (s *LocationService) GetList(companyId int) ([]*models.Location, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/locations/",
		new(models.Location),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.Location
	err = converter.SliceI2SliceModel(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *LocationService) Get(id int) (*models.Location, error) {
	model := new(models.Location)
	if err := s.client.getSourceSingle(id, "/companies/@all/locations/", model); err != nil {
		return nil, err
	}
	return model, nil
}
