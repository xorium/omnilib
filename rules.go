package omnimlib

type RuleData struct {
	ID    int    `jsonapi:"primary,rules"`
	Title string `jsonapi:"attr,title"`
	Slug  string `jsonapi:"attr,slug"`
	//Expression ??
	Duration          int    `jsonapi:"attr,duration"`
	EventLevel        string `jsonapi:"attr,eventLevel"`
	EventSessionState string `jsonapi:"attr,eventSessionState"`
	RuleGroup         string `jsonapi:"attr,ruleGroup"`
}

type RuleRelation struct {
	Company *CompanyData  `jsonapi:"relation,company"`
	Devices []*DeviceData `jsonapi:"relation,devices"`
	//params
}

type Rule struct {
	Data      *RuleData
	Relations *RuleRelation
}

type RuleService struct {
	client *Client
}
