package omnimlib

import "strconv"

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

type EventsSessionService struct {
	client *Client
}

func (s *EventsSessionService) GetList(companyId int) ([]EventsSession, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/events-sessions/",
		&Source{Data: new(EventsSessionData), Relations: new(EventsSessionRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []EventsSession
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *EventsSessionService) Get(id int) (*EventsSession, error) {
	data := new(EventsSessionData)
	rel := new(EventsSessionRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/events-sessions/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &EventsSession{data, rel}, nil
}
