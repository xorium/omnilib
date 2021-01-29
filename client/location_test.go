package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataLocation = map[string]string{
	"/companies/2/locations/": `
{
  "meta": {},
  "links": {},
  "data": [
    {
      "id": "1",
      "type": "locations",
      "attributes": {
        "name": "Sespel",
        "desc": "desc",
        "timezone": "+03:00",
        "info": {
          "infokey": "infoval"
        }
      },
      "relationships": {
        "company": {
          "data": {
            "id": "2",
            "type": "companies"
          }
        },
        "users": {
          "data": [
            {
              "id": "1",
              "type": "users"
            }
          ]
        },
        "children": {
          "data": [
            {
              "id": "3",
              "type": "locations"
            }
          ]
        }
      }
    },
    {
      "id": "2",
      "type": "locations",
      "attributes": {
        "name": "Sespel",
        "desc": "desc",
        "timezone": "+03:00",
        "info": {
          "infokey": "infoval"
        }
      },
      "relationships": {
        "company": {
          "data": {
            "id": "2",
            "type": "companies"
          }
        },
        "users": {
          "data": [
            {
              "id": "1",
              "type": "users"
            }
          ]
        },
        "children": {
          "data": [
            {
              "id": "3",
              "type": "locations"
            }
          ]
        }
      }
    }
  ],
  "included": [
    {
      "id": "2",
      "type": "companies",
      "attributes": {
        "name": "Sespel"
      }
    },
    {
      "id": "1",
      "type": "locations",
      "attributes": {
        "name": "Sespel",
        "desc": "desc",
        "timezone": "+03:00",
        "info": {
          "infokey": "infoval"
        }
      }
    },
    {
      "id": "1",
      "type": "users",
      "attributes": {
        "userName": "test",
        "password": "d404559f602eab6fd602ac7680dacbfaadd13630335e951f097af3900e9de176b6db28512f2e000b9d04fba5133e8b1c6e8df59db3a8ab9d60be4b97cc9e81db",
        "firstName": "Пользователь",
        "lastName": "Тестовый",
        "phoneNumber": "+79013801845",
        "email": "testemail",
        "image": "aaaa.jpg",
        "settings": {
          "aaaa": "sett1"
        }
      }
    }
  ]
}
`,
	///////////////////////////////////
	"/companies/@all/locations/1/": `
{
  "meta": {},
  "links": {},
  "data": 
    {
      "id": "1",
      "type": "locations",
      "attributes": {
        "name": "Sespel",
        "desc": "desc",
        "timezone": "+03:00",
        "info": {"infokey": "infoval"}
      },
      "relationships": {
        "company": {
          "data": {
            "id": "2",
            "type": "companies"
          }
        },
        "users": {
          "data": [{
            "id": "1",
            "type": "users"
          }]
        },
        "children": {
          "data": [
            {
              "id": "3",
              "type": "locations"
            }
          ]
        }
      }
    },
  "included": [
    {
      "id": "2",
      "type": "companies",
      "attributes": {
        "name": "Sespel"
      }
    },
    {
      "id": "1",
      "type": "locations",
      "attributes": {
        "name": "Sespel",
        "desc": "desc",
        "timezone": "+03:00",
        "info": {"infokey": "infoval"}
      }
    },
    {
      "id": "1",
      "type": "users",
      "attributes": {
        "userName": "test",
        "password": "d404559f602eab6fd602ac7680dacbfaadd13630335e951f097af3900e9de176b6db28512f2e000b9d04fba5133e8b1c6e8df59db3a8ab9d60be4b97cc9e81db",
        "firstName": "Пользователь",
        "lastName": "Тестовый",
        "phoneNumber": "+79013801845",
        "email": "testemail",
		"image": "aaaa.jpg",
        "settings": {
			"aaaa": "sett1"
		}
      }
    }
  ]
}
`,
}

func TestLocationService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	//c, err := NewClient(nil, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Location.GetList(2)
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

func TestLocationService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Location.Get(1)
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

func TestLocationService_GetNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServerNotFound))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	_, err = c.Location.Get(2)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

}
