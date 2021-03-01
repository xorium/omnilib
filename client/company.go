package client

import (
	models "gitlab.omnicube.ru/omnicube/omninanage/pkg/model/web"
	"omnilib/utils/converter"
)

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
	err = converter.SliceI2SliceModel(sources, &outSlice)
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
