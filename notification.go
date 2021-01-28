package omnimlib

type NotificationData struct {
	ID        int                    `jsonapi:"primary,users"`
	EventType string                 `jsonapi:"attr,eventType"`
	EventTime int                    `jsonapi:"attr,eventTime"`
	Info      map[string]interface{} `jsonapi:"attr,info"`
}

type NotificationRelation struct {
	Company      *CompanyData      `jsonapi:"relation,company"`
	Subscription *SubscriptionData `jsonapi:"relation,subscription"`
}

type Notification struct {
	Data      *NotificationData
	Relations *NotificationRelation
}

type NotificationService struct {
	client *Client
}
