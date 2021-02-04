package models

type Notification struct {
	ID           string                 `jsonapi:"primary,users"`
	EventType    string                 `jsonapi:"attr,eventType"`
	EventTime    int                    `jsonapi:"attr,eventTime"`
	Info         map[string]interface{} `jsonapi:"attr,info"`
	Company      *Company               `jsonapi:"relation,company"`
	Subscription *Subscription          `jsonapi:"relation,subscription"`
}
