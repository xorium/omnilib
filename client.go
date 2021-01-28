package omnimlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/jsonapi"
	gotils "github.com/savsgio/gotils/strconv"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"reflect"
	"strconv"
	"time"
)

type ClientConfig struct {
	BaseURL string
	TimeOut time.Duration

	Debug bool
}

type AuthConfig struct {
	WithAuthorization bool //if need authorization
	Username          string
	Password          string
}

type AuthData struct {
	AuthConfig
	token string
}

type Client struct {
	log      *zap.SugaredLogger
	config   ClientConfig
	authData AuthData

	TokenService  *TokenService
	Company       *CompanyService
	Device        *DeviceService
	DeviceGroup   *DeviceGroupService
	Parameter     *ParameterService
	User          *UserService
	Location      *LocationService
	Role          *RoleService
	Subscription  *SubscriptionService
	Manufacturer  *ManufacturerService
	DeviceModel   *DeviceModelService
	Event         *EventService
	EventsSession *EventsSessionService
}

type Source struct {
	Data      interface{}
	Relations interface{}
}

type ReqOptions struct {
	ContentType string
	Args        map[string]string
}

var DefaultConfig = ClientConfig{
	BaseURL: "https://omnimanage.omnicube.ru", //"http://172.26.1.80:8082",
	TimeOut: 5 * time.Second,

	Debug: false,
}

var DefaultAuth = AuthConfig{
	WithAuthorization: true,
	Username:          "ivan_p",
	Password:          "123",
}

//NewClient Create new client.
//conf - optional. If conf == nil => DefaultConfig
//auth - optional. If auth == nil => DefaultAuth
func NewClient(conf *ClientConfig, auth *AuthConfig) (*Client, error) {
	config := DefaultConfig
	if conf != nil {
		config = *conf
	}

	authConf := DefaultAuth
	if auth != nil {
		authConf = *auth
	}

	authData := AuthData{authConf, ""}

	logger, _ := zap.NewProduction()
	c := &Client{
		log:      logger.Sugar(),
		config:   config,
		authData: authData,
	}

	var err error
	c.TokenService = &TokenService{client: c}
	if c.IfNeedAuth() {
		c.authData.token, err = c.TokenService.GetNew()
		if err != nil {
			return nil, err
		}
	}

	c.Company = &CompanyService{client: c}
	c.Device = &DeviceService{client: c}
	c.DeviceGroup = &DeviceGroupService{client: c}
	c.Parameter = &ParameterService{client: c}
	c.User = &UserService{client: c}
	c.Location = &LocationService{client: c}
	c.Role = &RoleService{client: c}
	c.Subscription = &SubscriptionService{client: c}
	c.Manufacturer = &ManufacturerService{client: c}
	c.DeviceModel = &DeviceModelService{client: c}
	c.Event = &EventService{client: c}
	c.EventsSession = &EventsSessionService{client: c}
	return c, nil
}

func (c *Client) IfNeedAuth() bool {
	return c.authData.WithAuthorization
}

func (c *Client) getSourceSingle(id int, sourcePath string, val *Source) error {
	payload, err := c.doRequest("GET", sourcePath+strconv.Itoa(id)+"/", nil)
	if err != nil {
		return err
	}

	if val.Data != nil {
		err = jsonapi.UnmarshalPayload(bytes.NewReader(payload), val.Data)
		if err != nil {
			c.log.Errorf("Cant unmarshal payload: %v", err)
			return fmt.Errorf("Cant unmarshal payload %v", err)
		}
	}

	if val.Relations != nil {
		err = jsonapi.UnmarshalPayload(bytes.NewReader(payload), val.Relations)
		if err != nil {
			c.log.Errorf("Cant unmarshal payload: %v", err)
			return fmt.Errorf("Cant unmarshal payload %v", err)
		}
	}

	return nil
}

func (c *Client) getSourceMultiple(sourcePath string, sourceSingleRow *Source) ([]Source, error) {

	payload, err := c.doRequest("GET", sourcePath, nil)
	if err != nil {
		return nil, err
	}

	var recordsData []interface{}
	if sourceSingleRow.Data != nil {
		recordsData, err = jsonapi.UnmarshalManyPayload(bytes.NewReader(payload), reflect.TypeOf(sourceSingleRow.Data))
		if err != nil {
			c.log.Errorf("Cant unmarshal payload: %v", err)
			return nil, fmt.Errorf("Cant unmarshal payload: %v", err)
		}
	}

	var recordsRel []interface{}
	if sourceSingleRow.Relations != nil {
		recordsRel, err = jsonapi.UnmarshalManyPayload(bytes.NewReader(payload), reflect.TypeOf(sourceSingleRow.Relations))
		if err != nil {
			c.log.Errorf("Cant unmarshal payload: %v", err)
			return nil, fmt.Errorf("Cant unmarshal payload: %v", err)
		}
	}

	res := make([]Source, 0, 1)
	for i, rec := range recordsData {
		src := Source{}
		src.Data = rec
		if recordsRel != nil && len(recordsRel) >= i {
			src.Relations = recordsRel[i]
		}
		res = append(res, src)
	}

	return res, nil
}

func (c *Client) sourceSliceToOut(srcSlice []Source, out interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			c.log.Errorf("sourceSliceToOut panic: %v", r)
		}
	}()

	if reflect.TypeOf(out).Elem().Kind() != reflect.Slice {
		return fmt.Errorf("Out param is not a slice")
	}

	outRef := reflect.ValueOf(out)
	outSlice := reflect.MakeSlice(reflect.ValueOf(out).Elem().Type(), len(srcSlice), cap(srcSlice))

	for i, src := range srcSlice {
		outRow := outSlice.Index(i)

		if src.Data != nil {
			field := outRow.FieldByName("Data")
			if field.IsValid() {
				field.Set(reflect.ValueOf(src.Data))
			}
		}

		if src.Relations != nil {
			field := outRow.FieldByName("Relations")
			if field.IsValid() {
				field.Set(reflect.ValueOf(src.Relations))
			}
		}

	}
	outRef.Elem().Set(outSlice)

	return nil
}

func (c *Client) doRequest(method string, requestURI string, opt *ReqOptions) ([]byte, error) {
	//create request & response
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	//set headers
	req.Header.SetMethod(method)
	req.SetRequestURI(c.config.BaseURL + requestURI)

	if opt != nil && opt.ContentType != "" {
		req.Header.SetContentType(opt.ContentType)
	}
	if c.IfNeedAuth() {
		req.Header.Add("Authorization", c.authData.token)
	}

	if opt != nil && opt.Args != nil {
		for key, val := range opt.Args {
			req.PostArgs().Add(key, val)
		}
	}

	if c.config.Debug {
		req.Header.VisitAll(func(key, value []byte) {
			c.log.Debug(zap.String("key", gotils.B2S(key)), zap.String("value", gotils.B2S(value)))
		})
	}

	//send request
	err := fasthttp.DoTimeout(req, resp, c.config.TimeOut)
	if err != nil {
		c.log.Error("request error", err)
		if err == fasthttp.ErrTimeout {
			return nil, fmt.Errorf("timeout for %s", requestURI)
		}
		return nil, err
	}

	if c.config.Debug {
		c.log.Debugf("received response: %s", gotils.B2S(resp.Body()))
		//TODO add timings
	}

	//// try to parse error
	if resp.StatusCode() < fasthttp.StatusOK || resp.StatusCode() >= fasthttp.StatusBadRequest {
		c.log.Errorf("Errors: %v", resp.Body())
		errPayload := new(jsonapi.ErrorsPayload)
		if err = json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&errPayload); err == nil {
			if len(errPayload.Errors) > 0 {
				return nil, fmt.Errorf("code: %v,kkkkk title: %v, detail: %v ", errPayload.Errors[0].Code, errPayload.Errors[0].Title, errPayload.Errors[0].Detail)
			}
		}
		return nil, fmt.Errorf("Unknown Error, status code: %v", resp.StatusCode())
	}

	return resp.Body(), nil
}
