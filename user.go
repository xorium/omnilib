package omnimlib

import "strconv"

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
	Roles         []*RolesData        `jsonapi:"relation,roles"`
	Subscriptions []*SubscriptionData `jsonapi:"relation,subscriptions"`
}

type User struct {
	Data      *UserData
	Relations *UserRelation
}

type UserService struct {
	client *Client
}

func (s *UserService) GetList(companyId int) ([]User, error) {
	sources, err := s.client.getSourceMultiple(
		"/companies/"+strconv.Itoa(companyId)+"/users/",
		&Source{Data: new(UserData), Relations: new(UserRelation)},
	)
	if err != nil {
		return nil, err
	}

	var out []User
	err = s.client.sourceSliceToOut(sources, &out)

	return out, nil
}

func (s *UserService) Get(id int) (*User, error) {
	data := new(UserData)
	rel := new(UserRelation)
	if err := s.client.getSourceSingle(id, "/companies/@all/users/", &Source{Data: data, Relations: rel}); err != nil {
		return nil, err
	}
	return &User{data, rel}, nil
}
