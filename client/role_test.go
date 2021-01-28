package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataRole = map[string]string{
	"/companies/5/roles/": `
{
  "meta": {},
  "links": {},
  "data": [
    {
      "id": "1",
      "type": "roles",
      "attributes": {
        "name": "Адиминстратор",
        "slug": "admin",
        "assigned": true,
        "persistent": true,
		"info": {"info_key": "info_val"}
      },
      "relationships": {
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        }
      }
    },
    {
      "id": "2",
      "type": "roles",
      "attributes": {
        "name": "Кроссовые",
        "slug": "cross\n",
        "assigned": true,
        "persistent": true,
		"info": {"info_key": "info_val"}
      },
      "relationships": {
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        }
      }
    }
  ],
  "included": [
    {
      "id": "5",
      "type": "companies",
      "attributes": {
        "name": "Penta Hotels"
      }
    }
  ]
}
`,
	//////////////////////////
	"/companies/@all/roles/1/": `
{
  "meta": {},
  "links": {},
  "data": 
    {
      "id": "1",
      "type": "roles",
      "attributes": {
        "name": "Адиминстратор",
        "slug": "admin",
        "assigned": true,
        "persistent": true,
		"info": {"info_key": "info_val"}
      },
      "relationships": {
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        }
      }
    },
  "included": [
    {
      "id": "5",
      "type": "companies",
      "attributes": {
        "name": "Penta Hotels"
      }
    }
  ]
}
`,
}

func TestRoleService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	//c, err := NewClient(nil, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Role.GetList(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	t.Logf("\nresult: %#v", rec)

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

		err = IfHasEmptyField(v.Relations)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
	}

}

func TestRoleService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Role.Get(1)
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

	err = IfHasEmptyField(rec.Relations)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
}
