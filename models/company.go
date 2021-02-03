package models

type Company struct {
	ID   int    `jsonapi:"primary,companies"`
	Name string `jsonapi:"attr,name"`
}
