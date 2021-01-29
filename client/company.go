package client

import "omnilib/models"

type CompanyService struct {
	client *Client
}

func (s *CompanyService) GetList() ([]models.Company, error) {

	sources, err := s.client.getSourceMultiple("/companies/",
		&Source{Data: new(models.CompanyData)},
	)
	if err != nil {
		return nil, err
	}

	var outSlice []models.Company
	err = s.client.sourceSliceToOut(sources, &outSlice)

	return outSlice, nil
}

func (s *CompanyService) Get(id int) (*models.Company, error) {
	comp := new(models.CompanyData)

	if err := s.client.getSourceSingle(id, "/companies/", &Source{Data: comp}); err != nil {
		return nil, err
	}
	return &models.Company{Data: comp}, nil
}
