package omnimlib

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataCompany = map[string]string{
	"/companies/": `
{
  "meta": {},
  "links": {},
  "data": [
    {
      "id": "2",
      "type": "companies",
      "attributes": {
        "name": "Sespel"
      }
    },
    {
      "id": "4",
      "type": "companies",
      "attributes": {
        "name": "BetBoom"
      }
    },
    {
      "id": "1",
      "type": "companies",
      "attributes": {
        "name": "Omnicube"
      }
    }
  ],
  "included": []
}
`,
	"/companies/2/": `
{
			"meta": {},
			"links": {},
			"data": {
			"id": "2",
				"type": "companies",
				"attributes": {
				"name": "Sespel"
			}
		},
			"included": []
		}
`,
}

func TestCompanyService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Company.GetList()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if len(rec) != 3 {
		t.Errorf("wrong lines count in result, expected 3, got %v \n result: %#v", len(rec), rec)
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

func TestCompanyService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Company.Get(2)
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

func TestCompanyService_GetNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServerNotFound))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	_, err = c.Company.Get(2)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

}
