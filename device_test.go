package omnimlib

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var TestDataDevice = map[string]string{
	"/companies/2/devices/": `
{
  "meta": {},
  "links": {},
  "data": [
    {
    "id": "2",
    "type": "devices",
    "attributes": {
      "name": "M-710iC50",
      "slug": "slug",
      "title": "Fanuc M-710iC50",
      "desc": "Робот сварочный",
      "kind": "parent",
      "info": {
        "toolsNum": 1,
        "operatorsNum": 2,
        "connection": {
          "ip": "192.168.17.3",
          "port": 5000,
          "connecting_interval": 5,
          "ping_timeout_sec": 4
        }
      },
      "image": "http://omnicube.ru/images/devices-M710.png"
    },
    "relationships": {
      "company": {
        "data": {
          "id": "2",
          "type": "companies"
        }
      },
      "model": {
        "data": {
          "id": "1103",
          "type": "deviceModels"
        }
      },
      "groups": {
        "data": [
          {
            "id": "1045",
            "type": "deviceGroups"
          },
          {
            "id": "1043",
            "type": "deviceGroups"
          }
        ]
      },
      "location": {
        "data": {
          "id": "1086",
          "type": "locations"
        }
      },
      "parent": {
        "data": {
          "id": "999",
          "type": "devices",
          "meta": {}
        },
        "links": {},
        "meta": {}
      },
	  "rules": {
        "data": [
          {
            "id": "1",
            "type": "rules"
          },
          {
            "id": "2",
            "type": "rules"
          }
        ]
      }
    }
  },
{
    "id": "3",
    "type": "devices",
    "attributes": {
      "name": "2222",
      "slug": "slug22",
      "title": "Fanuc M-710iC50",
      "desc": "Робот сварочныйaaaaa",
      "kind": "parent",
      "info": {
        "toolsNum": 1,
        "operatorsNum": 2,
        "connection": {
          "ip": "192.168.17.3",
          "port": 5000,
          "connecting_interval": 5,
          "ping_timeout_sec": 4
        }
      },
      "image": "http://omnicube.ru/images/devices-M710.png"
    },
    "relationships": {
      "company": {
        "data": {
          "id": "2",
          "type": "companies"
        }
      },
      "model": {
        "data": {
          "id": "1103",
          "type": "deviceModels"
        }
      },
      "groups": {
        "data": [
          {
            "id": "1045",
            "type": "deviceGroups"
          },
          {
            "id": "1043",
            "type": "deviceGroups"
          }
        ]
      },
      "location": {
        "data": {
          "id": "1086",
          "type": "locations"
        }
      },
      "parent": {
        "data": {
          "id": "999",
          "type": "devices",
          "meta": {}
        },
        "links": {},
        "meta": {}
      },
	  "rules": {
        "data": [
          {
            "id": "1",
            "type": "rules"
          },
          {
            "id": "2",
            "type": "rules"
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
		"id": "999",
		"type": "devices",
		"attributes": {
		  "name": "PARENT",
		  "slug": "parent_slug",
		  "title": "parent_Fanuc M-710iC50",
		  "desc": "parent_Робот сварочный",
		  "kind": "parent",
		  "info": {
			"toolsNum": 1,
			"operatorsNum": 2,
			"connection": {
			  "ip": "192.168.17.3",
			  "port": 5000,
			  "connecting_interval": 5,
			  "ping_timeout_sec": 4
			}
		  },
		  "image": "http://omnicube.ru/images/devices-M710.png"
		}
	},
    {
      "id": "1103",
      "type": "deviceModels",
      "attributes": {
        "name": "M-710iC50",
        "title": "Fanuc M-710iC50",
        "desc": "Робот сварочный"
      }
    },
    {
      "id": "1086",
      "type": "locations",
      "attributes": {
        "name": "Цех 4",
        "timezone": "+03:00"
      }
    },
    {
      "id": "1045",
      "type": "deviceGroups",
      "attributes": {
        "name": "Сварочная группа",
        "type": "BUSINESS",
        "filters": {
          "operation": null,
          "operands": []
        }
      }
    },
    {
      "id": "1043",
      "type": "deviceGroups",
      "attributes": {
        "name": "Механическая обработка",
        "type": "BUSINESS",
        "filters": {
          "operation": null,
          "operands": []
        }
      }
    },
	{
		"id": "1",
		"type": "rules",
		"meta": {},
		"attributes": {
		  "title": "rule_title_1",
		  "slug": "battery",
		  "expression": {
			"+": "param_1"
		  },
		  "duration": 11,
		  "eventLevel": "eventLevel",
		  "eventSessionState": "sessionState",
		  "ruleGroup": "none",
		  "meta": [
			"meta",
			"info"
		  ]
		}
	},
	{
		"id": "2",
		"type": "rules",
		"meta": {},
		"attributes": {
		  "title": "rule_title_2",
		  "slug": "battery",
		  "expression": {
			"+": "param_1"
		  },
		  "duration": 11,
		  "eventLevel": "eventLevel",
		  "eventSessionState": "sessionState",
		  "ruleGroup": "none",
		  "meta": [
			"meta",
			"info"
		  ]
		}
	}
  ]
}
`,
	/////////////////////////////////////////////////////
	"/companies/@all/devices/2/": `
{
  "meta": {},
  "links": {},
  "data": {
    "id": "2",
    "type": "devices",
    "attributes": {
      "name": "M-710iC50",
      "slug": "bbb",
      "title": "Fanuc M-710iC50",
      "desc": "Робот сварочный",
      "kind": "parent",
      "info": {
        "toolsNum": 1,
        "operatorsNum": 2,
        "connection": {
          "ip": "192.168.17.3",
          "port": 5000,
          "connecting_interval": 5,
          "ping_timeout_sec": 4
        }
      },
      "image": "http://omnicube.ru/images/devices-M710.png"
    },
    "relationships": {
      "company": {
        "data": {
          "id": "2",
          "type": "companies"
        }
      },
      "model": {
        "data": {
          "id": "1103",
          "type": "deviceModels"
        }
      },
      "groups": {
        "data": [
          {
            "id": "1045",
            "type": "deviceGroups"
          },
          {
            "id": "1043",
            "type": "deviceGroups"
          }
        ]
      },
      "location": {
        "data": {
          "id": "1086",
          "type": "locations"
        }
      },
      "parent": {
        "data": {
          "id": "999",
          "type": "devices",
          "meta": {}
        },
        "links": {},
        "meta": {}
      },
	  "rules": {
        "data": [
          {
            "id": "1",
            "type": "rules"
          },
          {
            "id": "2",
            "type": "rules"
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
      "id": "1103",
      "type": "deviceModels",
      "attributes": {
        "name": "M-710iC50",
        "title": "Fanuc M-710iC50",
        "desc": "Робот сварочный"
      }
    },
    {
      "id": "1086",
      "type": "locations",
      "attributes": {
        "name": "Цех 4",
        "timezone": "+03:00"
      }
    },
    {
      "id": "1045",
      "type": "deviceGroups",
      "attributes": {
        "name": "Сварочная группа",
        "type": "BUSINESS",
        "filters": {
          "operation": null,
          "operands": []
        }
      }
    },
    {
      "id": "1043",
      "type": "deviceGroups",
      "attributes": {
        "name": "Механическая обработка",
        "type": "BUSINESS",
        "filters": {
          "operation": null,
          "operands": []
        }
      }
    },
	{
		"id": "999",
		"type": "devices",
		"attributes": {
		  "name": "PARENT",
		  "slug": "parent_slug",
		  "title": "parent_Fanuc M-710iC50",
		  "desc": "parent_Робот сварочный",
		  "kind": "parent",
		  "info": {
			"toolsNum": 1,
			"operatorsNum": 2,
			"connection": {
			  "ip": "192.168.17.3",
			  "port": 5000,
			  "connecting_interval": 5,
			  "ping_timeout_sec": 4
			}
		  },
		  "image": "http://omnicube.ru/images/devices-M710.png"
		}
	},
	{
		"id": "1",
		"type": "rules",
		"meta": {},
		"attributes": {
		  "title": "rule_title_1",
		  "slug": "battery",
		  "expression": {
			"+": "param_1"
		  },
		  "duration": 11,
		  "eventLevel": "eventLevel",
		  "eventSessionState": "sessionState",
		  "ruleGroup": "none",
		  "meta": [
			"meta",
			"info"
		  ]
		}
	},
	{
		"id": "2",
		"type": "rules",
		"meta": {},
		"attributes": {
		  "title": "rule_title_2",
		  "slug": "battery",
		  "expression": {
			"+": "param_1"
		  },
		  "duration": 11,
		  "eventLevel": "eventLevel",
		  "eventSessionState": "sessionState",
		  "ruleGroup": "none",
		  "meta": [
			"meta",
			"info"
		  ]
		}
	}
  ]
}
`,
}

func TestDeviceService_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Device.Get(2)
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

func TestDeviceService_GetList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	rec, err := c.Device.GetList(2)
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
	}
}

func TestDeviceService_GetNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServerNotFound))
	defer ts.Close()

	c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	_, err = c.Device.Get(2)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

}
