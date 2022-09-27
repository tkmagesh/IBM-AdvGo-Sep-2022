package models

//go:generate go run ../col-gen.go -N Customer -P models
type Customer struct {
	Id   int
	Name string
	City string
}
