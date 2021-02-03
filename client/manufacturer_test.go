package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataManufacturer = map[string]string{
	"/manufacturers/": `
{
  "meta": {},
  "links": {},
  "data": [
    {
      "id": "1",
      "type": "manufacturers",
      "attributes": {
        "name": "Fanuc",
      	"info": {"info_key": "info_val"}
		}
    },
    {
      "id": "2",
      "type": "manufacturers",
      "attributes": {
        "name": "ООО 'КР Пром'",
		"info": {"info_key": "info_val"}
      }
    }
  ],
  "included": []
}
`,
	"/manufacturers/1/": `
{
  "meta": {},
  "links": {},
  "data": {
    "id": "2",
    "type": "manufacturers",
    "attributes": {
      "name": "ООО 'КР Пром'",
      "info": {
        "info_key": "info_val"
      }
    }
  },
  "included": []
}
`,
}

func TestManufacturerService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Manufacturer.GetList()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if len(rec) != 2 {
		t.Errorf("wrong lines count in result, expected 2, got %v \n result: %#v", len(rec), rec)
		return
	}

	for _, v := range rec {
		err = IfHasEmptyField(v)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
	}

	t.Logf("\nresult: %#v", rec)
}

func TestManufacturerService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Manufacturer.Get(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	t.Logf("\nresult: %#v", rec)

	err = IfHasEmptyField(rec)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
}
