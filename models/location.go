package models

type LocationData struct {
	ID       int                    `jsonapi:"primary,locations"`
	Name     string                 `jsonapi:"attr,name"`
	Timezone string                 `jsonapi:"attr,timezone"`
	Info     map[string]interface{} `jsonapi:"attr,info"`
}

type LocationRelation struct {
	Company  *CompanyData    `jsonapi:"relation,company"`
	Children []*LocationData `jsonapi:"relation,children"`
	Users    []*UserData     `jsonapi:"relation,users"`
}

type Location struct {
	Data      *LocationData
	Relations *LocationRelation
}
