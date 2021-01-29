package models

type UserData struct {
	ID          int                    `jsonapi:"primary,users"`
	Name        string                 `jsonapi:"attr,userName"`
	Password    string                 `jsonapi:"attr,firstName"`
	FirstName   string                 `jsonapi:"attr,userName"`
	LastName    string                 `jsonapi:"attr,lastName"`
	PhoneNumber string                 `jsonapi:"attr,phoneNumber"`
	Email       string                 `jsonapi:"attr,email"`
	Image       string                 `jsonapi:"attr,image"`
	Settings    map[string]interface{} `jsonapi:"attr,settings"`
}

type UserRelation struct {
	Company       *CompanyData        `jsonapi:"relation,company"`
	Location      *LocationData       `jsonapi:"relation,location"`
	Roles         []*RoleData         `jsonapi:"relation,roles"`
	Subscriptions []*SubscriptionData `jsonapi:"relation,subscriptions"`
}

type User struct {
	Data      *UserData
	Relations *UserRelation
}
