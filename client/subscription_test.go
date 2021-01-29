package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataSubscription = map[string]string{
	"/companies/5/subscriptions/": `
{
  "meta": {},
  "links": {},
  "data": [
    {
      "id": "1",
      "type": "subscriptions",
      "attributes": {
        "title": "Тестовая подписка 1",
        "contactChannels": {
          "contact": "channels"
        },
        "options": {
          "options": "options"
        }
      },
      "relationships": {
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        },
        "user": {
          "data": {
            "id": "1",
            "type": "users"
          }
        },
        "rules": {
          "data": [
            {
              "id": "1",
              "type": "rules"
            }
          ]
        }
      }
    },
    {
      "id": "1",
      "type": "subscriptions",
      "attributes": {
        "title": "Тестовая подписка 1",
        "contactChannels": {
          "contact": "channels"
        },
        "options": {
          "options": "options"
        }
      },
      "relationships": {
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        },
        "user": {
          "data": {
            "id": "1",
            "type": "users"
          }
        },
        "rules": {
          "data": [
            {
              "id": "1",
              "type": "rules"
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
      "type": "users",
      "attributes": {
        "userName": "testov",
        "password": "d404559f602eab6fd602ac7680dacbfaadd13630335e951f097af3900e9de176b6db28512f2e000b9d04fba5133e8b1c6e8df59db3a8ab9d60be4b97cc9e81db",
        "firstName": "Тест",
        "lastName": "Тестов",
        "email": "testov-tt@omni.ru",
        "phoneNumber": "+79160000000"
      }
    },
    {
      "id": "1",
      "type": "rules",
      "attributes": {
        "title": "Нажата тревожная кнопка",
        "slug": "button_alarm",
        "expression": {
          "operation": "OR",
          "operands": [
            {
              "variable": "state",
              "operation": "==",
              "constant": "ALARM_ON"
            },
            {
              "variable": "state",
              "operation": "==",
              "constant": "ALARMING"
            }
          ]
        },
        "eventLevel": "ALARM",
        "eventSessionState": "open",
        "ruleGroup": "lora_button_1_alert"
      }
    }
  ]
}
`,
	"/companies/@all/subscriptions/1/": `
{
  "meta": {},
  "links": {},
  "data": {
    "id": "1",
    "type": "subscriptions",
    "attributes": {
      "title": "Тестовая подписка 1",
      "contactChannels": {
        "contact": "channels"
      },
      "options": {
        "options": "options"
      }
    },
    "relationships": {
      "company": {
        "data": {
          "id": "5",
          "type": "companies"
        }
      },
      "user": {
        "data": {
          "id": "1",
          "type": "users"
        }
      },
      "rules": {
        "data": [
          {
            "id": "1",
            "type": "rules"
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
      "type": "users",
      "attributes": {
        "userName": "testov",
        "password": "d404559f602eab6fd602ac7680dacbfaadd13630335e951f097af3900e9de176b6db28512f2e000b9d04fba5133e8b1c6e8df59db3a8ab9d60be4b97cc9e81db",
        "firstName": "Тест",
        "lastName": "Тестов",
        "email": "testov-tt@omni.ru",
        "phoneNumber": "+79160000000"
      }
    },
    {
      "id": "1",
      "type": "rules",
      "attributes": {
        "title": "Нажата тревожная кнопка",
        "slug": "button_alarm",
        "expression": {
          "operation": "OR",
          "operands": [
            {
              "variable": "state",
              "operation": "==",
              "constant": "ALARM_ON"
            },
            {
              "variable": "state",
              "operation": "==",
              "constant": "ALARMING"
            }
          ]
        },
        "eventLevel": "ALARM",
        "eventSessionState": "open",
        "ruleGroup": "lora_button_1_alert"
      }
    }
  ]
}
`,
}

func TestSubscriptionService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	//c, err := NewClient(nil, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Subscription.GetList(5)
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

func TestSubscriptionService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Subscription.Get(1)
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
