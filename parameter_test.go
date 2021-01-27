package omnimlib

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataParameter = map[string]string{
	"/parameters/": `
	{
	"meta": {},
	  "links": {
		"self": null,
		"related": null
	  },
	  "data": [
		{
		  "id": "1",
		  "type": "parameters",
		  "meta": {},
		  "attributes": {
			"name": "laser_temp",
			"desc": "descr1",
			"type": "float",
			"isValuesSetFinite": true,
			"info": {"info1": "info111"}
		  }
		},
		{
		  "id": "2",
		  "type": "parameters",
		  "meta": {},
		  "attributes": {
			"name": "device_kpi:useful_percent",
			"desc": "descr2",
			"type": "float",
			"isValuesSetFinite": true,
			"info": {"info1": "info111"}
		  }
		}
	],
  	"included": []
	}
`,

	"/parameters/1/": `
	{
	  "meta": {},
	  "links": {
		"self": null,
		"related": null
	  },
	  "data": {
		"id": "1",
		"type": "parameters",
		"meta": {},
		"attributes": {
		  "name": "laser_temp",
		  "desc": "descr1",
		  "type": "float",
		  "isValuesSetFinite": true,
		  "info": {"info1": "info111"}
		}
	  },
	  "included": []
	}
`,
}

func TestParameterService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Parameter.GetList()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if len(rec) != 2 {
		t.Errorf("wrong lines count in result, expected 2, got %v \n result: %#v", len(rec), rec)
		return
	}

	for _, v := range rec {
		err = IfHasEmptyField(v.Data)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
	}

	t.Logf("\nresult: %#v", rec)
}

func TestParameterService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Parameter.Get(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	t.Logf("\nresult: %#v", rec)

	err = IfHasEmptyField(rec.Data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
}

func TestParameterService_GetNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServerNotFound))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	_, err = c.Parameter.Get(2)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

}
