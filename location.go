package omnimlib

type LocationData struct {
	ID       int    `jsonapi:"primary,locations"`
	Name     string `jsonapi:"attr,name"`
	Timezone string `jsonapi:"attr,timezone"`
}
