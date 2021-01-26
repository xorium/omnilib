# omnilib
Go library for working with OmniManage.

This library is still a work in progress

##Functionality Supported
**Devices**
- Get One
- Get list

**Device group**
- Get One
- Get list

**Company**
- Get One
- Get list

**Token**
- Get New

**Other**
...

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


