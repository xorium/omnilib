package omnimlib

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const ValidToken = "12345"

func OmniServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	var payload string
	switch {
	case r.RequestURI == "/tokens/":
		payload = `
{
  "meta": {},
  "links": {
    "self": null,
    "related": null
  },
  "data": {
    "id": "7357",
    "type": "tokens",
    "meta": {},
    "attributes": {
      "value": "12345"
    }
  },
  "included": []
}
`
	case r.RequestURI == "/companies/":
		payload = `
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
`

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(payload))
}

func TestCompanyService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))

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

	t.Logf("\nresult: %#v", rec)
}

func TestCompanyService_Get(t *testing.T) {
	c, err := NewClient(nil, &AuthConfig{WithAuthorization: false})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Company.Get(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	t.Logf("\nresult: %#v", rec)
}
