# omnilib
Go library for working with OmniManage.

This library is still a work in progress

##Functionality Supported
**Company**
- Get One
- Get list

**Parameter**
- Get One
- Get list

**Device**
- Get One
- Get list

**Device group**
- Get One
- Get list

**Token**
- Get New

**Location**
- Get One
- Get list

**Rule**
only struct

**Device-model**
only struct

**Manufacturer**
none

**Role**
only struct

**Websocket channel**
none

**User**
- Get One
- Get list

**Notification**
none

**Converter**
none

**Subscription**
only struct

**Event sessions**
none

**Event**
none


##Example
Get Device list:
```go
cli, err := NewClient(nil, nil)
if err != nil {
//..
}
devices, err := c.Device.GetList(company_id)
if err != nil {
//..
}
```
Get single device by Id:
```go
cli, err := NewClient(nil, nil)
if err != nil {
//..
}
device, err := c.Device.Get(device_id)
if err != nil {
//..
}
```


