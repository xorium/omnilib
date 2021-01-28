package omnimlib

import "strconv"

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

type Event struct {
	Data      *EventData
	Relations *EventRelation
}

type EventService struct {
	client *Client
}

func (s *EventService) GetList(companyId int) ([]Event, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/events/",
		&Source{Data: new(EventData), Relations: new(EventRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []Event
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *EventService) Get(id int) (*Event, error) {
	data := new(EventData)
	rel := new(EventRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/events/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &Event{data, rel}, nil
}
