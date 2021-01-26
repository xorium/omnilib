package omnimlib

import (
	"fmt"
	"testing"
)

func TestDeviceGroupService_Get(t *testing.T) {
	c, _ := NewClient(nil, nil)
	rec, err := c.DeviceGroup.Get(4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	fmt.Printf("value: %v", rec)
}

func TestDeviceGroupService_GetList(t *testing.T) {
	c, _ := NewClient(nil, nil)

	recs, err := c.DeviceGroup.GetList(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	fmt.Printf("value: %v", recs)
}
