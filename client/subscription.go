package client

import (
	"omnilib/models"
	"strconv"
)

type SubscriptionService struct {
	client *Client
}

func (s *SubscriptionService) GetList(companyId int) ([]models.Subscription, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/subscriptions/",
		&Source{Data: new(models.SubscriptionData), Relations: new(models.SubscriptionRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []models.Subscription
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *SubscriptionService) Get(id int) (*models.Subscription, error) {
	data := new(models.SubscriptionData)
	rel := new(models.SubscriptionRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/subscriptions/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &models.Subscription{data, rel}, nil
}
