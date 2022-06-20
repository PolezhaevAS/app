package models

type Service struct {
	ID      uint64
	Name    string
	Methods []*Method
}

type Method struct {
	ID   uint64
	Name string
}
