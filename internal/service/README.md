### Develop

## Add new service

First step is add file ${SERVICE_NAME}.go

Then put into file your service
```
func NewYourServiceName() *Service {
	service := new(Service)
	service.Methods = make(map[string]bool)

	service.Name = "YourServiceName"
    service.Methods["MethodNameInProtoFile"] = true | false 

	return service
}
```

True or false - this check to needed JWT token 

## Example
```
func NewAccessService() *Service {
	service := new(Service)
	service.Methods = make(map[string]bool)

	service.Name = "access"
	service.Methods["UserAccess"] = false
	service.Methods["List"] = true

	return service
}
```

In service.go need put your new service in New func.

Example:
```
    ...

    case "access":
		return NewAccessService(), nil

    ...
```

And then add line in All func:
```
...

    services = append(services, NewAccessService())
    
...
```
