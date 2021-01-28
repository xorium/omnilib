package models

type DeviceData struct {
	ID          int                    `jsonapi:"primary,devices"`
	Name        string                 `jsonapi:"attr,name"`
	Slug        string                 `jsonapi:"attr,slug"`
	Title       string                 `jsonapi:"attr,title"`
	Description string                 `jsonapi:"attr,desc"`
	Kind        string                 `jsonapi:"attr,kind"`
	Info        map[string]interface{} `jsonapi:"attr,info"`
	Image       string                 `jsonapi:"attr,image"`

	//Info        struct {
	//	ToolsNum     int `jsonapi:"attr,toolsNum"`
	//	OperatorsNum int `jsonapi:"attr,operatorsNum"`
	//	Connection   struct {
	//		IP                 string `jsonapi:"attr,ip"`
	//		Port               int    `jsonapi:"attr,port"`
	//		ConnectingInterval int    `jsonapi:"attr,connecting_interval"`
	//		PingTimeoutSec     int    `jsonapi:"attr,ping_timeout_sec"`
	//		RecvLengthByte     int    `jsonapi:"attr,recv_length_byte"`
	//		RecvTerminator     string `jsonapi:"attr,recv_terminator"`
	//	} `jsonapi:"attr,connection"`
	//} `jsonapi:"attr,info"`
}

type DeviceRelation struct {
	Company  *CompanyData       `jsonapi:"relation,company"`
	Model    *DeviceModelData   `jsonapi:"relation,model"`
	Location *LocationData      `jsonapi:"relation,location"`
	Groups   []*DeviceGroupData `jsonapi:"relation,groups"`
	Parent   *DeviceData        `jsonapi:"relation,parent"`
	Rules    []*RuleData        `jsonapi:"relation,rules"`
}

type Device struct {
	Data      *DeviceData
	Relations *DeviceRelation
}

type DeviceGroupData struct {
	ID          int                    `jsonapi:"primary,deviceGroups"`
	Name        string                 `jsonapi:"attr,name"`
	Description string                 `jsonapi:"attr,desc"`
	Type        string                 `jsonapi:"attr,type"`
	Filters     map[string]interface{} `jsonapi:"attr,filters"`
}

type DeviceGroupRelation struct {
	Company *CompanyData  `jsonapi:"relation,company"`
	Devices []*DeviceData `jsonapi:"relation,devices"`
	User    *UserData     `jsonapi:"relation,user"`
}

type DeviceGroup struct {
	Data      *DeviceGroupData
	Relations *DeviceGroupRelation
}
