package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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
	case r.RequestURI == "/companies/" || r.RequestURI == "/companies/2/":
		payload = TestDataCompany[r.RequestURI]

	case r.RequestURI == "/companies/2/devices/" || r.RequestURI == "/companies/@all/devices/2/":
		payload = TestDataDevice[r.RequestURI]

	case r.RequestURI == "/companies/2/device-groups/" || r.RequestURI == "/companies/@all/device-groups/4/":
		payload = TestDataDeviceGroup[r.RequestURI]

	case r.RequestURI == "/parameters/" || r.RequestURI == "/parameters/1/":
		payload = TestDataParameter[r.RequestURI]

	case r.RequestURI == "/companies/5/users/" || r.RequestURI == "/companies/@all/users/1/":
		payload = TestDataUser[r.RequestURI]

	case r.RequestURI == "/companies/2/locations/" || r.RequestURI == "/companies/@all/locations/1/":
		payload = TestDataLocation[r.RequestURI]

	case r.RequestURI == "/companies/5/roles/" || r.RequestURI == "/companies/@all/roles/1/":
		payload = TestDataRole[r.RequestURI]

	case r.RequestURI == "/companies/5/subscriptions/" || r.RequestURI == "/companies/@all/subscriptions/1/":
		payload = TestDataSubscription[r.RequestURI]

	case r.RequestURI == "/manufacturers/" || r.RequestURI == "/manufacturers/1/":
		payload = TestDataManufacturer[r.RequestURI]

	case r.RequestURI == "/companies/5/events-sessions/" || r.RequestURI == "/companies/@all/events-sessions/1/":
		payload = TestDataEventsSession[r.RequestURI]

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(payload))
}

func OmniServerNotFound(w http.ResponseWriter, r *http.Request) {

	var payload string
	switch {
	case r.RequestURI == "/tokens/":
		w.WriteHeader(http.StatusOK)
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
	default:
		w.WriteHeader(http.StatusNotFound)
		payload = `
{"errors":[{"id":"1611667533.2231216","status":"404","code":"RESOURCE_ERROR","title":"RESOURCE_OBJECT_NOT_FOUND","detail":"Resource: resource_name, id"}]}
`
	}

	w.Write([]byte(payload))
}

func OmniServerNoAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)

	payload := `
{
  "errors": [
    {
      "id": "1611667938.1506572",
      "status": "401",
      "code": "RESOURCE_ERROR",
      "title": "AUTH_REQUIRED",
      "detail": "???????????????????? ????????????????????????????"
    }
  ]
}
`

	w.Write([]byte(payload))
}

func IfHasEmptyField(v interface{}) error {
	s := reflect.ValueOf(v).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		v := reflect.ValueOf(f.Interface())
		if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
			return fmt.Errorf("%v is empty", s.Type().Field(i).Name)
		}
	}
	return nil
}

func TestClient_NoAuth(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(OmniServerNoAuth))
	defer ts.Close()

	_, err := NewClient(&Config{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
