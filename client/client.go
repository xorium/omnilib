package client

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

type Config struct {
	BaseURL string
	TimeOut time.Duration

	Debug bool
}

type AuthConfig struct {
	WithAuthorization bool //if need authorization
	Username          string
	Password          string
	JWT               string
}

type Client struct {
	log        *zap.SugaredLogger
	config     Config
	authConfig AuthConfig

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

type ReqOptions struct {
	ContentType string
	Args        map[string]string
}

var DefaultConfig = Config{
	BaseURL: "https://omnimanage.omnicube.ru", //"http://172.26.1.80:8082",
	TimeOut: 5 * time.Second,

	Debug: false,
}

var DefaultAuth = AuthConfig{
	WithAuthorization: true,
	Username:          "omnilib",
	Password:          "netcubE1881",
}

// NewClient Create new client.
// conf - optional. If conf == nil => DefaultConfig
// auth - optional. If auth == nil => DefaultAuth
func NewClient(conf *Config, auth *AuthConfig) (*Client, error) {
	config := DefaultConfig
	if conf != nil {
		config = *conf
	}

	authConf := DefaultAuth
	if auth != nil {
		authConf = *auth
	}

	logger, _ := zap.NewProduction()
	c := &Client{
		log:        logger.Sugar(),
		config:     config,
		authConfig: authConf,
	}

	var err error
	c.TokenService = &TokenService{client: c}
	if c.IfNeedAuth() {
		c.authConfig.JWT, err = c.TokenService.GetNew()
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
	return c.authConfig.JWT == ""
}

func (c *Client) getSourceSingle(id int, sourcePath string, model interface{}) error {
	payload, err := c.doRequest("GET", sourcePath+strconv.Itoa(id)+"/", nil)
	if err != nil {
		return err
	}

	err = jsonapi.UnmarshalPayload(bytes.NewReader(payload), model)
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil
}

func (c *Client) getSourceMultiple(sourcePath string, modelSingleRow interface{}) ([]interface{}, error) {

	payload, err := c.doRequest("GET", sourcePath, nil)
	if err != nil {
		return nil, err
	}

	records, err := jsonapi.UnmarshalManyPayload(bytes.NewReader(payload), reflect.TypeOf(modelSingleRow))
	if err != nil {
		c.log.Error(err)
		return nil, err
	}
	return records, nil
}

func (c *Client) doRequest(method string, requestURI string, opt *ReqOptions) ([]byte, error) {
	// create request & response
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// set headers
	req.Header.SetMethod(method)
	req.SetRequestURI(c.config.BaseURL + requestURI)

	if opt != nil && opt.ContentType != "" {
		req.Header.SetContentType(opt.ContentType)
	}
	if c.IfNeedAuth() {
		req.Header.Add("Authorization", c.authConfig.JWT)
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

	// send request
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
		// TODO add timings
	}

	// try to parse error
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
