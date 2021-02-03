package client

import (
	"bytes"
	"fmt"
	"github.com/google/jsonapi"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type Country struct {
	ID   string `jsonapi:"primary,countries"`
	Name string `jsonapi:"attr,name"`
}

type Company struct {
	ID      string   `jsonapi:"primary,companies"`
	Name    string   `jsonapi:"attr,name"`
	Country *Country `jsonapi:"relation,country"`
}

type Role struct {
	ID         string                 `jsonapi:"primary,roles"`
	Name       string                 `jsonapi:"attr,name"`
	Slug       string                 `jsonapi:"attr,slug"`
	Assigned   bool                   `jsonapi:"attr,assigned"`
	Persistent bool                   `jsonapi:"attr,persistent"`
	Info       map[string]interface{} `jsonapi:"attr,info"`
	Company    *Company               `jsonapi:"relation,company"`
}

func TestRoleExample(t *testing.T) {
	resp := `
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
	      },
	      "relationships": {
	        "country": {
	          "data": {
	            "id": "2",
	            "type": "countries"
	          }
	        }
	      }
	    }
	  ]
	}
	`
	role := new(Role)
	err := jsonapi.UnmarshalPayload(strings.NewReader(resp), role)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "1", role.ID)
	assert.Equal(t, "5", role.Company.ID)
	assert.Equal(t, "2", role.Company.Country.ID)

	buf := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalPayload(buf, role); err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))
}
