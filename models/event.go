package models

type EventData struct {
	ID           int                    `jsonapi:"primary,events"`
	Type         string                 `jsonapi:"attr,type"`
	Title        string                 `jsonapi:"attr,title"`
	Time         int                    `jsonapi:"attr,time"`
	SessionId    string                 `jsonapi:"attr,sessionId"`
	SessionSlug  string                 `jsonapi:"attr,sessionSlug"`
	SessionState string                 `jsonapi:"attr,sessionState"`
	Level        string                 `jsonapi:"attr,level"`
	Ttl          int                    `jsonapi:"attr,ttl"`
	Info         map[string]interface{} `jsonapi:"attr,info"`
}

type EventRelation struct {
	Company  *CompanyData   `jsonapi:"relation,company"`
	Location *LocationData  `jsonapi:"relation,location"`
	Device   *LocationData  `jsonapi:"relation,device"`
	User     *UserData      `jsonapi:"relation,user"`
	Session  *EventsSession `jsonapi:"relation,session"`
}

type EventsSessionData struct {
	ID            string `jsonapi:"primary,eventsSessions"`
	Title         string `jsonapi:"attr,title"`
	State         string `jsonapi:"attr,state"`
	Level         string `jsonapi:"attr,level"`
	LastEventTime int    `jsonapi:"attr,lastEventTime"`
	Slug          string `jsonapi:"attr,slug"`
}

type EventsSessionRelation struct {
	Company  *CompanyData  `jsonapi:"relation,company"`
	Device   *DeviceData   `jsonapi:"relation,device"`
	Location *LocationData `jsonapi:"relation,location"`
	LastUser *UserData     `jsonapi:"relation,lastUser"`
	Events   []*EventData  `jsonapi:"relation,events"`
}

type EventsSession struct {
	Data      *EventsSessionData
	Relations *EventsSessionRelation
}
