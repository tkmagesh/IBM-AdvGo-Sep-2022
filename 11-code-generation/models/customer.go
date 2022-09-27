package models

//go:generate go run ../col-gen.go -N Customer -P models
//go:generate go fmt Customers.go
type Customer struct {
	Id   int
	Name string
	City string
}
