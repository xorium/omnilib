# omnilib
Go library for working with OmniManage.

This library is still a work in progress

##Functionality Supported##
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
only struct (no data in omimanage)

**Device-model**
no data in omnimanage => no tests

**Manufacturer**
- Get One
- Get list

**Role**
- Get One
- Get list

**Websocket channel**
//??????

**User**
- Get One
- Get list

**Notification**
only struct (???), no data

**Conversion**
//??????

**Subscription**
- Get One
- Get list

**Event**
no data in omnimanage => no tests

**Event sessions**
- Get One
- Get list


##Example
Get Device list:
```go
cli, err := NewClient(nil, nil) //nil params == default settings
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
cli, err := NewClient(nil, nil) //nil params == default settings
if err != nil {
//..
}
device, err := c.Device.Get(device_id)
if err != nil {
//..
}
```


