package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataEventsSession = map[string]string{
	"/companies/5/events-sessions/": `
{
  "meta": {},
  "links": {},
  "data": [
    {
      "id": "a300e0b4-0ac0-42d2-8366-29bc37fe6c2c",
      "type": "eventsSessions",
      "attributes": {
        "state": "closed",
        "title": "Нажата тревожная кнопка",
        "level": "ALARM",
        "slug": "button_alarm",
        "lastEventTime": 1611324042
      },
      "relationships": {
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        },
        "device": {
          "data": {
            "id": "1",
            "type": "devices"
          }
        },
        "lastUser": {
          "data": {
            "id": "1",
            "type": "users"
          }
        },
        "location": {
          "data": {
            "id": "1",
            "type": "locations"
          }
        },
        "events": {
          "data": [
            {
              "id": "1",
              "type": "events"
            },
            {
              "id": "2",
              "type": "events"
            }
          ]
        }
      }
    },
    {
      "id": "bbbbe0b4-0ac0-42d2-8366-29bc37fe6c2c",
      "type": "eventsSessions",
      "attributes": {
        "state": "closed",
        "title": "Нажата тревожная кнопка",
        "level": "ALARM",
        "slug": "button_alarm",
        "lastEventTime": 1611324042
      },
      "relationships": {
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        },
        "device": {
          "data": {
            "id": "1",
            "type": "devices"
          }
        },
        "lastUser": {
          "data": {
            "id": "1",
            "type": "users"
          }
        },
        "location": {
          "data": {
            "id": "1",
            "type": "locations"
          }
        },
        "events": {
          "data": [
            {
              "id": "1",
              "type": "events"
            },
            {
              "id": "2",
              "type": "events"
            }
          ]
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
    },
    {
      "id": "1",
      "type": "events",
      "attributes": {
        "type": "rules:trigger",
        "title": "Нажата тревожная кнопка",
        "time": 1611323964,
        "level": "ALARM",
        "info": {
          "param_name": "state",
          "param_value": "ALARM_ON",
          "company_name": "Penta Hotels",
          "rule_title": "Нажата тревожная кнопка",
          "userId": "1658",
          "userName": "omnilib"
        },
        "sessionId": "a300e0b4-0ac0-42d2-8366-29bc37fe6c2c",
        "sessionSlug": "button_alarm",
        "sessionState": "open"
      }
    },
    {
      "id": "2",
      "type": "events",
      "attributes": {
        "type": "rules:trigger",
        "title": "Нажата тревожная кнопка",
        "time": 1611324042,
        "level": "ALARM",
        "info": {
          "param_name": "state",
          "param_value": "ALARM_OFF",
          "company_name": "Penta Hotels",
          "rule_title": "Нажата тревожная кнопка",
          "userId": "1658",
          "userName": "omnilib"
        },
        "sessionId": "a300e0b4-0ac0-42d2-8366-29bc37fe6c2c",
        "sessionSlug": "button_alarm",
        "sessionState": "closed"
      }
    },
    {
      "id": "1",
      "type": "devices",
      "attributes": {
        "name": "lora_button_11",
        "slug": "3832333854307C02",
        "desc": "Номер 11",
        "title": "Тревожная кнопка",
        "image": "http://server.local/none.jpg",
        "kind": "parent",
        "info": {
          "deveui": "3832333854307C02",
          "meta_id": "11"
        }
      }
    },
    {
      "id": "1",
      "type": "locations",
      "attributes": {
        "name": "room_821",
        "desc": "Комната 821",
        "timezone": "+03:00",
        "info": {
          "type": "room",
          "room": "821",
          "number": "21",
          "floor": "8"
        }
      }
    },
    {
      "id": "1",
      "type": "users",
      "attributes": {
        "userName": "omnilib",
        "password": "735892fb7bd37ca62c42a64d7d33bd2124533da1ab9932d8f0f988587fdd4bb84dac5c0f9d4b3289b64f5b68c9971e338f1c4ea1edc26b6264501c97c247186e",
        "firstName": "Omnilib",
        "email": "developers@netcube.ru"
      }
    }
  ]
}
`,
	"/companies/@all/events-sessions/1/": `
{
  "meta": {},
  "links": {},
  "data": {
    "id": "a300e0b4-0ac0-42d2-8366-29bc37fe6c2c",
    "type": "eventsSessions",
    "attributes": {
      "state": "closed",
      "title": "Нажата тревожная кнопка",
      "level": "ALARM",
      "slug": "button_alarm",
      "lastEventTime": 1611324042
    },
    "relationships": {
      "company": {
        "data": {
          "id": "5",
          "type": "companies"
        }
      },
      "device": {
        "data": {
          "id": "1",
          "type": "devices"
        }
      },
      "lastUser": {
        "data": {
          "id": "1",
          "type": "users"
        }
      },
      "location": {
        "data": {
          "id": "1",
          "type": "locations"
        }
      },
      "events": {
        "data": [
          {
            "id": "1",
            "type": "events"
          },
          {
            "id": "2",
            "type": "events"
          }
        ]
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
    },
    {
      "id": "1",
      "type": "events",
      "attributes": {
        "type": "rules:trigger",
        "title": "Нажата тревожная кнопка",
        "time": 1611323964,
        "level": "ALARM",
        "info": {
          "param_name": "state",
          "param_value": "ALARM_ON",
          "company_name": "Penta Hotels",
          "rule_title": "Нажата тревожная кнопка",
          "userId": "1658",
          "userName": "omnilib"
        },
        "sessionId": "a300e0b4-0ac0-42d2-8366-29bc37fe6c2c",
        "sessionSlug": "button_alarm",
        "sessionState": "open"
      }
    },
    {
      "id": "2",
      "type": "events",
      "attributes": {
        "type": "rules:trigger",
        "title": "Нажата тревожная кнопка",
        "time": 1611324042,
        "level": "ALARM",
        "info": {
          "param_name": "state",
          "param_value": "ALARM_OFF",
          "company_name": "Penta Hotels",
          "rule_title": "Нажата тревожная кнопка",
          "userId": "1658",
          "userName": "omnilib"
        },
        "sessionId": "a300e0b4-0ac0-42d2-8366-29bc37fe6c2c",
        "sessionSlug": "button_alarm",
        "sessionState": "closed"
      }
    },
    {
      "id": "1",
      "type": "devices",
      "attributes": {
        "name": "lora_button_11",
        "slug": "3832333854307C02",
        "desc": "Номер 11",
        "title": "Тревожная кнопка",
        "image": "http://server.local/none.jpg",
        "kind": "parent",
        "info": {
          "deveui": "3832333854307C02",
          "meta_id": "11"
        }
      }
    },
    {
      "id": "1",
      "type": "locations",
      "attributes": {
        "name": "room_821",
        "desc": "Комната 821",
        "timezone": "+03:00",
        "info": {
          "type": "room",
          "room": "821",
          "number": "21",
          "floor": "8"
        }
      }
    },
    {
      "id": "1",
      "type": "users",
      "attributes": {
        "userName": "omnilib",
        "password": "735892fb7bd37ca62c42a64d7d33bd2124533da1ab9932d8f0f988587fdd4bb84dac5c0f9d4b3289b64f5b68c9971e338f1c4ea1edc26b6264501c97c247186e",
        "firstName": "Omnilib",
        "email": "developers@netcube.ru"
      }
    }
  ]
}
`,
}

func TestEventSessionService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.EventsSession.GetList(5)
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
		err = IfHasEmptyField(v)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
	}

}

func TestEventsSessionService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.EventsSession.Get(1)
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
