package omnimlib

type RolesData struct {
	ID         int    `jsonapi:"primary,roles"`
	Name       string `jsonapi:"attr,name"`
	Slug       string `jsonapi:"attr,slug"`
	Assigned   bool   `jsonapi:"attr,assigned"`
	Persistent bool   `jsonapi:"attr,persistent"`
	//Info??
}
