package client

import "strconv"

type SubscriptionData struct {
	ID              int                    `jsonapi:"primary,subscriptions"`
	Title           string                 `jsonapi:"attr,title"`
	ContactChannels map[string]interface{} `jsonapi:"attr,contactChannels"`
	Options         map[string]interface{} `jsonapi:"attr,options"`
}

type SubscriptionRelation struct {
	Company *CompanyData `jsonapi:"relation,company"`
	User    *UserData    `jsonapi:"relation,user"`
	Rules   []*RuleData  `jsonapi:"relation,rules"`
}

type Subscription struct {
	Data      *SubscriptionData
	Relations *SubscriptionRelation
}

type SubscriptionService struct {
	client *Client
}

func (s *SubscriptionService) GetList(companyId int) ([]Subscription, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/subscriptions/",
		&Source{Data: new(SubscriptionData), Relations: new(SubscriptionRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []Subscription
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *SubscriptionService) Get(id int) (*Subscription, error) {
	data := new(SubscriptionData)
	rel := new(SubscriptionRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/subscriptions/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &Subscription{data, rel}, nil
}
