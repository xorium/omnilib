package client

import (
	"gitlab.omnicube.ru/libs/omnilib/models"
	"strconv"
)

type SubscriptionService struct {
	client *Client
}

func (s *SubscriptionService) GetList(companyId int) ([]*models.Subscription, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/subscriptions/",
		new(models.Subscription),
	)
	if err != nil {
		return nil, err
	}

	var outSlice []*models.Subscription
	err = s.client.sourceSliceToOut(sources, &outSlice)
	if err != nil {
		return nil, err
	}
	return outSlice, nil
}

func (s *SubscriptionService) Get(id int) (*models.Subscription, error) {
	model := new(models.Subscription)
	if err := s.client.getSourceSingle(id, "/companies/@all/subscriptions/", model); err != nil {
		return nil, err
	}
	return model, nil
}
