package models

type CompanyData struct {
	ID   int    `jsonapi:"primary,companies"`
	Name string `jsonapi:"attr,name"`
}

type Company struct {
	Data *CompanyData

	//*CompanyData
}
