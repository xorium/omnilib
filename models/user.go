package models

type User struct {
	ID            string                 `jsonapi:"primary,users"`
	Name          string                 `jsonapi:"attr,userName"`
	Password      string                 `jsonapi:"attr,firstName"`
	FirstName     string                 `jsonapi:"attr,userName"`
	LastName      string                 `jsonapi:"attr,lastName"`
	PhoneNumber   string                 `jsonapi:"attr,phoneNumber"`
	Email         string                 `jsonapi:"attr,email"`
	Image         string                 `jsonapi:"attr,image"`
	Settings      map[string]interface{} `jsonapi:"attr,settings"`
	Company       *Company               `jsonapi:"relation,company"`
	Location      *Location              `jsonapi:"relation,location"`
	Roles         []*Role                `jsonapi:"relation,roles"`
	Subscriptions []*Subscription        `jsonapi:"relation,subscriptions"`
}
