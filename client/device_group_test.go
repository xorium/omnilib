package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataDeviceGroup = map[string]string{
	"/companies/2/device-groups/": `
{
  "meta": {},
  "links": {},
  "data": [
    {
      "id": "4",
      "type": "deviceGroups",
      "attributes": {
        "name": "Сварочная группа",
        "desc": "aaa",
        "type": "BUSINESS",
        "filters": {
          "aaaa": "bbbbb"
        }
      },
      "relationships": {
        "company": {
          "data": {
            "id": "2",
            "type": "companies"
          }
        },
        "user": {
          "data": {
            "id": "1",
            "type": "users",
            "meta": {}
          },
          "links": {},
          "meta": {}
        },
        "devices": {
          "data": [
            {
              "id": "1",
              "type": "devices"
            },
            {
              "id": "2",
              "type": "devices"
            }
          ]
        }
      }
    },
    {
      "id": "5",
      "type": "deviceGroups",
      "attributes": {
        "name": "СТП сварочная группа",
        "desc": "bbbb",
        "type": "BUSINESS",
        "filters": {
          "aaaa": "bbbbb"
        }
      },
      "relationships": {
        "company": {
          "data": {
            "id": "2",
            "type": "companies"
          }
        },
        "user": {
          "data": {
            "id": "1",
            "type": "users",
            "meta": {}
          },
          "links": {},
          "meta": {}
        },
        "devices": {
          "data": [
            {
              "id": "1",
              "type": "devices"
            },
            {
              "id": "3",
              "type": "devices"
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
      "type": "devices",
      "attributes": {
        "name": "STP",
        "desc": "Станция сварки трением с перемешиванием",
        "title": "СТП (Саша)",
        "image": "http://omnicube.ru/images/devices-stp.png",
        "kind": "parent",
        "info": {
          "toolsNum": 1,
          "operatorsNum": 2,
          "connection": {
            "ip": "192.168.18.210",
            "port": 9696,
            "connecting_interval": 5,
            "ping_timeout_sec": 4,
            "recv_length_byte": 1024,
            "recv_terminator": "\n"
          }
        }
      }
    },
    {
      "id": "2",
      "type": "devices",
      "attributes": {
        "name": "STP_4PL",
        "desc": "Станция сварки трением с перемешиванием",
        "title": "СТП 4PL",
        "image": "http://omnicube.ru/images/devices-stp4pl.png",
        "kind": "parent",
        "info": {
          "toolsNum": 1,
          "operatorsNum": 2,
          "connection": {
            "ip": "192.168.18.220",
            "port": 9696,
            "connecting_interval": 5,
            "ping_timeout_sec": 4,
            "recv_length_byte": 1024,
            "recv_terminator": "\n"
          }
        }
      }
    },
    {
      "id": "3",
      "type": "devices",
      "attributes": {
        "name": "ARCMate_100iC7L2",
        "desc": "Робот сварочный",
        "title": "Fanuc Robot Mate 100iC/7L (МАКС 2)",
        "image": "http://omnicube.ru/images/devices-AM100.png",
        "kind": "parent",
        "info": {
          "toolsNum": 1,
          "operatorsNum": 2,
          "connection": {
            "ip": "192.168.17.43",
            "port": 5000,
            "connecting_interval": 5,
            "ping_timeout_sec": 4
          }
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
        "settings": {}
      }
    }
  ]
}
`,

	/////////////////////////////////////////
	"/companies/@all/device-groups/4/": `
{
  "meta": {},
  "links": {},
  "data": {
    "id": "4",
    "type": "deviceGroups",
    "meta": {},
    "attributes": {
      "name": "device_group_name",
      "desc": "description",
      "type": "SERVICE",
      "filters": {
        "locations": [
          2,
          3,
          6
        ],
        "groups": [
          "welder"
        ],
        "devices": [
          1,
          2
        ],
        "manufacturers": [
          "fanuc"
        ]
      }
    },
    "relationships": {
      "company": {
        "data": {
          "id": "2",
          "type": "companies",
          "meta": {}
        },
        "links": {},
        "meta": {}
      },
      "user": {
        "data": {
          "id": "1",
          "type": "users",
          "meta": {}
        },
        "links": {},
        "meta": {}
      },
      "devices": {
        "data": [
          {
            "id": "1",
            "type": "devices"
          },
          {
            "id": "3",
            "type": "devices"
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
      "type": "devices",
      "attributes": {
        "name": "STP",
        "desc": "Станция сварки трением с перемешиванием",
        "title": "СТП (Саша)",
        "image": "http://omnicube.ru/images/devices-stp.png",
        "kind": "parent",
        "info": {
          "toolsNum": 1,
          "operatorsNum": 2,
          "connection": {
            "ip": "192.168.18.210",
            "port": 9696,
            "connecting_interval": 5,
            "ping_timeout_sec": 4,
            "recv_length_byte": 1024,
            "recv_terminator": "\n"
          }
        }
      }
    },
    {
      "id": "2",
      "type": "devices",
      "attributes": {
        "name": "STP_4PL",
        "desc": "Станция сварки трением с перемешиванием",
        "title": "СТП 4PL",
        "image": "http://omnicube.ru/images/devices-stp4pl.png",
        "kind": "parent",
        "info": {
          "toolsNum": 1,
          "operatorsNum": 2,
          "connection": {
            "ip": "192.168.18.220",
            "port": 9696,
            "connecting_interval": 5,
            "ping_timeout_sec": 4,
            "recv_length_byte": 1024,
            "recv_terminator": "\n"
          }
        }
      }
    },
    {
      "id": "3",
      "type": "devices",
      "attributes": {
        "name": "ARCMate_100iC7L2",
        "desc": "Робот сварочный",
        "title": "Fanuc Robot Mate 100iC/7L (МАКС 2)",
        "image": "http://omnicube.ru/images/devices-AM100.png",
        "kind": "parent",
        "info": {
          "toolsNum": 1,
          "operatorsNum": 2,
          "connection": {
            "ip": "192.168.17.43",
            "port": 5000,
            "connecting_interval": 5,
            "ping_timeout_sec": 4
          }
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
			"settings": {}
      		}
    	}
  ]
}
`,
}

func TestDeviceGroupService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.DeviceGroup.Get(4)
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

func TestDeviceGroupService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.DeviceGroup.GetList(2)
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
