package client

import (
	"omnilib/models"
	"strconv"
)

type LocationService struct {
	client *Client
}

func (s *LocationService) GetList(companyId int) ([]models.Location, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/locations/",
		&Source{Data: new(models.LocationData), Relations: new(models.LocationRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []models.Location
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *LocationService) Get(id int) (*models.Location, error) {
	data := new(models.LocationData)
	rel := new(models.LocationRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/locations/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &models.Location{data, rel}, nil
}
