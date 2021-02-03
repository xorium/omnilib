package client

import "omnilib/models"

type CompanyService struct {
	client *Client
}

func (s *CompanyService) GetList() ([]*models.Company, error) {

	sources, err := s.client.getSourceMultiple("/companies/",
		new(models.Company),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.Company
	err = s.client.sourceSliceToOut(sources, &outSlice)
	if err != nil {
		return nil, err
	}

	return outSlice, nil
}

func (s *CompanyService) Get(id int) (*models.Company, error) {
	model := new(models.Company)

	if err := s.client.getSourceSingle(id, "/companies/", model); err != nil {
		return nil, err
	}
	return model, nil
}
