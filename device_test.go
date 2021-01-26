package omnimlib

import (
	"fmt"
	"testing"
)

func TestDeviceService_Get(t *testing.T) {
	c, _ := NewClient(nil, nil)
	rec, err := c.Device.Get(4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	fmt.Printf("value: %v", rec)
}

func TestDeviceService_GetList(t *testing.T) {
	c, _ := NewClient(nil, nil)

	recs, err := c.Device.GetList(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	fmt.Printf("value: %v", recs)
}
