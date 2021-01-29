package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataUser = map[string]string{
	"/companies/5/users/": `
	{
  "meta": {},
  "links": {},
  "data": [
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
      },
      "relationships": {
		"location": {
          "data": {
            "id": "1",
            "type": "locations"
          }
        },
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        },
        "roles": {
          "data": [
            {
              "id": "1",
              "type": "roles"
            },
            {
              "id": "2",
              "type": "roles"
            }
          ]
        },
        "subscriptions": {
          "data": [
            {
              "id": "1",
              "type": "subscriptions"
            },
            {
              "id": "2",
              "type": "subscriptions"
            }
          ]
        }
      }
    },
    {
      "id": "2",
      "type": "users",
      "attributes": {
        "userName": "user1",
        "password": "012c2c42e65db51d828d2550f49b837e483bff650028696c3d7962bffd535159139eb2f7e849dd6676310ca882c3befff57adb992de8b29adaf00d1380c1e0f0",
        "firstName": "Алексей",
        "lastName": "Шабанов",
        "image": "http://omnicube.ru/images/xxx.jpg",
        "email": "test email2",
		"phoneNumber": "+79013801222",
		"image": "aaaa.jpg",
        "settings": {
			"aaaa": "sett1"
		}
      },
      "relationships": {
		"location": {
          "data": {
            "id": "1",
            "type": "locations"
          }
        },
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        },
        "roles": {
          "data": [
            {
              "id": "1",
              "type": "roles"
            },
            {
              "id": "2",
              "type": "roles"
            }
          ]
        },
        "subscriptions": {
          "data": [
            {
              "id": "1",
              "type": "subscriptions"
            },
            {
              "id": "2",
              "type": "subscriptions"
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
      "type": "locations",
      "attributes": {
        "name": "Цех 4",
        "timezone": "+03:00"
      }
    },
    {
      "id": "1",
      "type": "subscriptions",
      "attributes": {
        "title": "Тревожная кнопка",
        "contactChannels": {
          "telegram": "ON"
        }
      }
    },
    {
      "id": "2",
      "type": "subscriptions",
      "attributes": {
        "title": "Датчик температуры",
        "contactChannels": {
          "telegram": "ON"
        }
      }
    },
    {
      "id": "1",
      "type": "roles",
      "attributes": {
        "name": "Кроссовые",
        "slug": "cross\n",
        "persistent": true
      }
    },
    {
      "id": "2",
      "type": "roles",
      "attributes": {
        "name": "Жилые",
        "slug": "rooms",
        "persistent": true
      }
    }
  ]
}
`,
	//////////////////////////
	"/companies/@all/users/1/": `
	{
  "meta": {},
  "links": {},
  "data": 
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
      },
      "relationships": {
		"location": {
          "data": {
            "id": "1",
            "type": "locations"
          }
        },
        "company": {
          "data": {
            "id": "5",
            "type": "companies"
          }
        },
        "roles": {
          "data": [
            {
              "id": "1",
              "type": "roles"
            },
            {
              "id": "2",
              "type": "roles"
            }
          ]
        },
        "subscriptions": {
          "data": [
            {
              "id": "1",
              "type": "subscriptions"
            },
            {
              "id": "2",
              "type": "subscriptions"
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
      "type": "locations",
      "attributes": {
        "name": "Цех 4",
        "timezone": "+03:00"
      }
    },
    {
      "id": "1",
      "type": "subscriptions",
      "attributes": {
        "title": "Тревожная кнопка",
        "contactChannels": {
          "telegram": "ON"
        }
      }
    },
    {
      "id": "2",
      "type": "subscriptions",
      "attributes": {
        "title": "Датчик температуры",
        "contactChannels": {
          "telegram": "ON"
        }
      }
    },
    {
      "id": "1",
      "type": "roles",
      "attributes": {
        "name": "Кроссовые",
        "slug": "cross\n",
        "persistent": true
      }
    },
    {
      "id": "2",
      "type": "roles",
      "attributes": {
        "name": "Жилые",
        "slug": "rooms",
        "persistent": true
      }
    }
  ]
}
`,
}

func TestUserService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.User.GetList(5)
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

func TestUserService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.User.Get(1)
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
