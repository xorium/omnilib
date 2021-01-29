package models

type RoleData struct {
	ID         int                    `jsonapi:"primary,roles"`
	Name       string                 `jsonapi:"attr,name"`
	Slug       string                 `jsonapi:"attr,slug"`
	Assigned   bool                   `jsonapi:"attr,assigned"`
	Persistent bool                   `jsonapi:"attr,persistent"`
	Info       map[string]interface{} `jsonapi:"attr,info"`
}

type RoleRelation struct {
	Company *CompanyData `jsonapi:"relation,company"`
}

type Role struct {
	Data      *RoleData
	Relations *RoleRelation
}
