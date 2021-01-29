package models

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
