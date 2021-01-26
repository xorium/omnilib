package omnimlib

type DeviceModelData struct {
	ID          int    `jsonapi:"primary,deviceModels"`
	Name        string `jsonapi:"attr,name"`
	Title       string `jsonapi:"attr,title"`
	Description string `jsonapi:"attr,desc"`
}
