package omnimlib

//func TestEventService_GetList(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(OmniServer))
//	defer ts.Close()
//
//	//c, err := NewClient(&ClientConfig{BaseURL: ts.URL, TimeOut: time.Second * 5}, nil)
//	c, err := NewClient(nil, nil)
//	if err != nil {
//		t.Errorf("Unexpected error: %v", err)
//		return
//	}
//
//	rec, err := c.Event.GetList(4)
//	if err != nil {
//		t.Errorf("Unexpected error: %v", err)
//		return
//	}
//	t.Logf("\nresult: %#v", rec)
//
//	if len(rec) != 2 {
//		t.Errorf("wrong lines count in result, expected 2, got %v \n result: %#v", len(rec), rec)
//		return
//	}
//
//	for _, v := range rec {
//		err = IfHasEmptyField(v.Data)
//		if err != nil {
//			t.Errorf("Unexpected error: %v", err)
//			return
//		}
//
//		err = IfHasEmptyField(v.Relations)
//		if err != nil {
//			t.Errorf("Unexpected error: %v", err)
//			return
//		}
//	}
//
//}
