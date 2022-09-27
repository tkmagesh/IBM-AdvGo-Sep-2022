package models

//go:generate go run ../col-gen.go -N Product -P models
type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}
