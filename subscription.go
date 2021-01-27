package omnimlib

type SubscriptionData struct {
	ID              int                    `jsonapi:"primary,subscriptions"`
	Title           string                 `jsonapi:"attr,title"`
	ContactChannels map[string]interface{} `jsonapi:"attr,contactChannels"`
	Options         map[string]interface{} `jsonapi:"attr,options"`
}
