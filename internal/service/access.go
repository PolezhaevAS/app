package service

func NewAccessService() *Service {
	service := new(Service)
	service.Methods = make(map[string]bool)

	service.Name = "access"
	service.Methods["UserAccesses"] = false
	service.Methods["List"] = true
	service.Methods["Group"] = true
	service.Methods["CreateGroup"] = true
	service.Methods["UpdateGroup"] = true
	service.Methods["DeleteGroup"] = true
	service.Methods["Users"] = true
	service.Methods["AddUser"] = true
	service.Methods["RemoveUser"] = true
	service.Methods["ListServices"] = true
	service.Methods["AddMethod"] = true
	service.Methods["RemoveMethod"] = true

	return service
}
