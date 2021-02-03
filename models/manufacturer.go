package models

type Manufacturer struct {
	ID   int                    `jsonapi:"primary,manufacturers"`
	Name string                 `jsonapi:"attr,name"`
	Info map[string]interface{} `jsonapi:"attr,info"`
}
